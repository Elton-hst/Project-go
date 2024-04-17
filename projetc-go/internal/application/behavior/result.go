package behavior

import "github.com/Elton-hst/internal/application/logger"

type Result struct {
    Success interface{}
    Problem error
}

func NewResult() *Result { return &Result{} }

func newSuccess(value interface{}) *Result {
    return &Result{Success: value}
}

func failure(err error) *Result {
    return &Result{Problem: err}
}

func (r *Result) ResultCheck(Result interface{}, err error) (*Result, error) {
    if err != nil {
        logger.Error.Println(err)
        return failure(err), err
    }

    return newSuccess(Result), nil
}

func (r *Result) ErrorCheck(err error) error {
    if err != nil {
        logger.Error.Println(err)
    }
    return err
}

func (r *Result) IsSuccess() bool {
    return r.Problem == nil
}

func (r *Result) IsFailure() bool {
    return r.Problem != nil
}

func (r *Result) GetSuccess() interface{} {
    return r.Success
}

func (r *Result) GetFailure() error {
    return r.Problem
}
