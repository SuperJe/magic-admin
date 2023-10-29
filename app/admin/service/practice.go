package service

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go-admin/app/admin/service/dto"
	"gorm.io/gorm"

	"github.com/SuperJe/coco/app/data_proxy/model"
)

type Practice struct {
	service.Service
}

func (p *Practice) GetPracticeCode(ctx context.Context, ids []int32, uid int) (*dto.GetPracticeCodeRsp, error) {
	if len(ids) == 0 || uid == 0 {
		return nil, fmt.Errorf("empty param")
	}
	cps := make([]*dto.CPPPractice, 0)
	if err := p.Orm.Where("user_id = ? AND p_id IN ?", uid, ids).Find(&cps).Error; err != nil {
		return nil, err
	}
	rsp := &dto.GetPracticeCodeRsp{
		BaseRsp: dto.BaseRsp{},
		Codes:   map[int64]string{},
	}
	for _, cp := range cps {
		cp := cp
		rsp.Codes[cp.ProblemID] = cp.Code
	}
	return rsp, nil
}

func (p *Practice) SavePracticeCode(ctx context.Context, uid, id int64, code string) error {
	if id == 0 || uid == 0 {
		return fmt.Errorf("empty param")
	}
	cond := &dto.CPPPractice{UserID: uid, ProblemID: id}
	cp := &dto.CPPPractice{}
	err := p.Orm.First(cp, cond).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		now := time.Now()
		cp = &dto.CPPPractice{ProblemID: id, UserID: uid, Code: code, Updated: now, Created: now}
		return p.Orm.Create(cp).Error
	}
	return p.Orm.Model(cp).UpdateColumn("code", code).Error
}

func (p *Practice) SubmitPracticeCode(ctx context.Context, uid, id int64, code string) (*dto.SubmitPracticeCodeRsp, error) {
	// 先写入DB
	if err := p.SavePracticeCode(ctx, uid, id, code); err != nil {
		return nil, err
	}

	path, _ := os.Getwd()
	path = fmt.Sprintf("%s/common/problem/practice/cpp/practice_cpp_%d.txt", path, id)
	// 打开文件（以只读模式）
	file, err := os.Open(path)
	if err != nil {
		// 处理错误
		return nil, errors.Wrapf(err, "open file %s err", path)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("close %s err: %s", path, err.Error())
		}
	}()

	rsp := &dto.SubmitPracticeCodeRsp{Accept: true}
	// 创建一个 Scanner 对象来从文件中读取内容
	scanner := bufio.NewScanner(file)
	// 逐行读取文件内容
	// 文件都是以3|2格式开头，表示接下来3行是输入，2行是输出。
	// 一直读取到文件末尾
	for scanner.Scan() {
		in, expOut, err := readPracticeData(scanner)
		if err != nil {
			return nil, err
		}

		actOut, err := getActualOutput(in, code)
		if err != nil {
			return nil, err
		}
		if expOut != actOut {
			return nil, fmt.Errorf("WRONG! 输入样例:\n%s\n\n您的输出:\n%s\n\n期望输出:\n%s\n", in, actOut, expOut)
		}
	}

	return rsp, nil
}

func getActualOutput(in, code string) (string, error) {
	reqBody := &model.RunCompilerReq{
		Lang:  model.LangCPP,
		Code:  code,
		Input: in,
	}
	body, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "http://127.0.0.1:7777/compiler/run", bytes.NewBuffer(body))
	if err != nil {
		return "", errors.Wrap(err, "http.NewRequest err")
	}
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "cli.Do err")
	}
	defer func() {
		if err := rsp.Body.Close(); err != nil {
			_ = rsp.Body.Close()
		}
	}()
	bs, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", errors.Wrap(err, "ReadAll err")
	}
	if rsp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http err code:%d", rsp.StatusCode)
	}
	data := &model.RunCompilerRsp{}
	if err := json.Unmarshal(bs, data); err != nil {
		return "", errors.Wrap(err, "unmarshal err")
	}
	fmt.Println("compiler rsp:", data)
	if data.Code != 0 {
		return "", fmt.Errorf("compiler err:%s", data.Msg)
	}
	return data.OutPut, nil
}

func readPracticeData(scanner *bufio.Scanner) (string, string, error) {
	line := scanner.Text() // 获取当前行的内容
	contents := strings.Split(line, "|")
	if len(contents) != 2 {
		return "", "", fmt.Errorf("content err:%s", line)
	}
	fmt.Println("line:", line) // 打印当前行的内容
	inputLines := cast.ToInt(contents[0])
	outputLines := cast.ToInt(contents[1])
	inputString := ""
	for i := 0; i < inputLines; i++ {
		if scanner.Scan() {
			inputString += scanner.Text()
			if i != inputLines-1 {
				inputString += "\n"
			}
		}
	}

	// 读取输出部分
	outputString := ""
	for i := 0; i < outputLines; i++ {
		if scanner.Scan() {
			outputString += scanner.Text()
		}
		if i != outputLines-1 {
			outputString += "\n"
		}
	}

	// 检查 Scanner 是否出现错误
	if err := scanner.Err(); err != nil {
		return "", "", fmt.Errorf("read file err: %s", err.Error())
	}
	return inputString, outputString, nil
}
