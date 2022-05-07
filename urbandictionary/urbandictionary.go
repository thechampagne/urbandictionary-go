package urbandictionary

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type response struct {
	List []struct {
		Definition  string    `json:"definition"`
		Permalink   string    `json:"permalink"`
		ThumbsUp    int       `json:"thumbs_up"`
		SoundUrls   []string  `json:"sound_urls"`
		Author      string    `json:"author"`
		Word        string    `json:"word"`
		Defid       int       `json:"defid"`
		CurrentVote string    `json:"current_vote"`
		WrittenOn   string `json:"written_on"`
		Example     string    `json:"example"`
		ThumbsDown  int       `json:"thumbs_down"`
	} `json:"list"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type Response struct {
	Definition  string
	Permalink   string
	ThumbsUp    int
	SoundUrls   []string
	Author      string
	Word        string
	Defid       int
	WrittenOn   string
	Example     string
	ThumbsDown  int
}

func get(endpoint string) (string, error) {
	response, err := http.Get(fmt.Sprintf("https://api.urbandictionary.com/v0/%s", endpoint))
	if err != nil {
		return "", errors.New("")
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("")
	}
	return string(body), nil
}

type urbanDictionary struct {
	term string
	page int32
}

func isError(s string) error {
	var data errorResponse
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		return errors.New("error")
	}
	if len(data.Error) != 0 {
		return errors.New(strings.ToLower(data.Error))
	}
	return nil
}

func New(input string, page int32) urbanDictionary {
	return urbanDictionary{input,page}
}

func (u urbanDictionary) Data() ([]Response, error) {
	resp, err := get(fmt.Sprintf("define?term=%s&page=%d", u.term, u.page))
	if err != nil {
		return []Response{}, newError("error")
	}
	isError := isError(resp)
	if isError != nil {
		return []Response{}, newError(isError.Error())
	}
	var response response
	jsonErr := json.Unmarshal([]byte(resp), &response)
	if jsonErr != nil {
		return []Response{}, newError("error")
	}
	if len(response.List) == 0 {
		return []Response{}, newError("empty data")
	}
	var responseSlice []Response
	for _, v := range response.List {
		responseSlice = append(responseSlice,Response{
			v.Definition,
			v.Permalink,
			v.ThumbsUp,
			v.SoundUrls,
			v.Author,
			v.Word,
			v.Defid,
			v.WrittenOn,
			v.Example,
			v.ThumbsDown,
		})
	}
	return responseSlice, nil
}

func Random() ([]Response, error) {
	resp, err := get("random")
	if err != nil {
		return []Response{}, newError("error")
	}
	isError := isError(resp)
	if isError != nil {
		return []Response{}, newError(isError.Error())
	}
	var response response
	jsonErr := json.Unmarshal([]byte(resp), &response)
	if jsonErr != nil {
		return []Response{}, newError("error")
	}
	if len(response.List) == 0 {
		return []Response{}, newError("empty data")
	}
	var responseSlice []Response
	for _, v := range response.List {
		responseSlice = append(responseSlice,Response{
			v.Definition,
			v.Permalink,
			v.ThumbsUp,
			v.SoundUrls,
			v.Author,
			v.Word,
			v.Defid,
			v.WrittenOn,
			v.Example,
			v.ThumbsDown,
		})
	}
	return responseSlice, nil
}

func DefinitionById(id int64) (Response, error) {
	resp, err := get(fmt.Sprintf("define?defid=%d", id))
	if err != nil {
		return Response{}, newError("error")
	}
	isError := isError(resp)
	if isError != nil {
		return Response{}, newError(isError.Error())
	}
	var response response
	jsonErr := json.Unmarshal([]byte(resp), &response)
	if jsonErr != nil {
		return Response{}, newError("error")
	}
	if len(response.List) == 0 {
		return Response{}, newError("empty data")
	}
	var responseSlice Response
	for _, v := range response.List {
		responseSlice = Response{
			v.Definition,
			v.Permalink,
			v.ThumbsUp,
			v.SoundUrls,
			v.Author,
			v.Word,
			v.Defid,
			v.WrittenOn,
			v.Example,
			v.ThumbsDown,
		}
	}
	return responseSlice, nil
}

func ToolTip(term string) (string, error) {
	resp, err := get(fmt.Sprintf("tooltip?term=%s", term))
	if err != nil {
		return "", newError("error")
	}
	isError := isError(resp)
	if isError != nil {
		return "", newError(isError.Error())
	}
	var response map[string]string
	jsonErr := json.Unmarshal([]byte(resp), &response)
	if jsonErr != nil {
		return "", newError("error")
	}
	if response["string"] ==  "" {
		return "", newError("error")
	}

	return response["string"], nil
}