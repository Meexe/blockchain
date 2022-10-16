// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson2ff71951DecodeGithubComMeexeBlockchainInternalModels(in *jlexer.Lexer, out *Block) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "hash":
			out.HashSum = string(in.String())
		case "prevHash":
			out.PrevHash = string(in.String())
		case "transactions":
			(out.Transactions).UnmarshalEasyJSON(in)
		case "ts":
			out.Ts = int64(in.Int64())
		case "pow":
			out.Pow = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2ff71951EncodeGithubComMeexeBlockchainInternalModels(out *jwriter.Writer, in Block) {
	out.RawByte('{')
	first := true
	_ = first
	if in.HashSum != "" {
		const prefix string = ",\"hash\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.HashSum))
	}
	{
		const prefix string = ",\"prevHash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PrevHash))
	}
	{
		const prefix string = ",\"transactions\":"
		out.RawString(prefix)
		(in.Transactions).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"ts\":"
		out.RawString(prefix)
		out.Int64(int64(in.Ts))
	}
	{
		const prefix string = ",\"pow\":"
		out.RawString(prefix)
		out.Int64(int64(in.Pow))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Block) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2ff71951EncodeGithubComMeexeBlockchainInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Block) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2ff71951EncodeGithubComMeexeBlockchainInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Block) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2ff71951DecodeGithubComMeexeBlockchainInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Block) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2ff71951DecodeGithubComMeexeBlockchainInternalModels(l, v)
}