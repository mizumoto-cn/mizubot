/*
 * Copyright (c) 2024 Ruiyuan "mizumoto-cn" Xu
 *
 * This file is part of "github.com/mizumoto-cn/mizubot".
 *
 * Licensed under the Mizumoto General Public License v1.5 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://github.com/mizumoto-cn/mizubot/blob/main/LICENSE
 *     https://github.com/mizumoto-cn/mizubot/blob/main/licensing
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package core

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mizumoto-cn/dailyreport/pkg/util"
)

func CreateReport(tPath string, cPath string) (string, error) {
	tLines, err := util.ReadFile(tPath)
	if err != nil {
		return "", errors.New("Failed to read template file")
	}
	cLines, err := util.ReadFile(cPath)
	if err != nil {
		return "", errors.New("Failed to read content file")
	}
	return createReport(tLines, cLines, NewDailyReportFormatter())
}

type Formatter func([]string, []string) (string, error)

func createReport(tLines []string, cLines []string, f Formatter) (string, error) {
	return f(tLines, cLines)
}

func NewDailyReportFormatter() Formatter {
	return func(tLines []string, cLines []string) (string, error) {
		// from ${{0}} to ${{n-1}} in tLines where n is the length of cLines
		// replace ${{i}} with cLines[i]
		// return the result
		var result strings.Builder
		cnt := 0
		for _, line := range tLines {
			for cnt < len(cLines) {
				if strings.Contains(line, "${{"+fmt.Sprint(cnt)+"}}") {
					line = strings.ReplaceAll(line, "${{"+fmt.Sprint(cnt)+"}}", cLines[cnt])
					cnt++
				} else {
					break
				}
			}
			result.WriteString(line)
			result.WriteString("\n")
		}
		if cnt < len(cLines)-1 {
			for i := cnt; i < len(cLines); i++ {
				result.WriteString(cLines[i])
				result.WriteString("\n")
			}
		}
		return result.String(), nil
	}
}
