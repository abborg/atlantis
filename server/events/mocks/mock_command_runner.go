// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: CommandRunner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	events "github.com/runatlantis/atlantis/server/events"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockCommandRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCommandRunner(options ...pegomock.Option) *MockCommandRunner {
	mock := &MockCommandRunner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockCommandRunner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockCommandRunner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockCommandRunner) RunAutoplanCommand(baseRepo models.Repo, headRepo models.Repo, pull models.PullRequest, user models.User) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandRunner().")
	}
	_params := []pegomock.Param{baseRepo, headRepo, pull, user}
	pegomock.GetGenericMockFrom(mock).Invoke("RunAutoplanCommand", _params, []reflect.Type{})
}

func (mock *MockCommandRunner) RunCommentCommand(baseRepo models.Repo, maybeHeadRepo *models.Repo, maybePull *models.PullRequest, user models.User, pullNum int, cmd *events.CommentCommand) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommandRunner().")
	}
	_params := []pegomock.Param{baseRepo, maybeHeadRepo, maybePull, user, pullNum, cmd}
	pegomock.GetGenericMockFrom(mock).Invoke("RunCommentCommand", _params, []reflect.Type{})
}

func (mock *MockCommandRunner) VerifyWasCalledOnce() *VerifierMockCommandRunner {
	return &VerifierMockCommandRunner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockCommandRunner) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockCommandRunner {
	return &VerifierMockCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockCommandRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockCommandRunner {
	return &VerifierMockCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockCommandRunner) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockCommandRunner {
	return &VerifierMockCommandRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockCommandRunner struct {
	mock                   *MockCommandRunner
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockCommandRunner) RunAutoplanCommand(baseRepo models.Repo, headRepo models.Repo, pull models.PullRequest, user models.User) *MockCommandRunner_RunAutoplanCommand_OngoingVerification {
	_params := []pegomock.Param{baseRepo, headRepo, pull, user}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RunAutoplanCommand", _params, verifier.timeout)
	return &MockCommandRunner_RunAutoplanCommand_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommandRunner_RunAutoplanCommand_OngoingVerification struct {
	mock              *MockCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommandRunner_RunAutoplanCommand_OngoingVerification) GetCapturedArguments() (models.Repo, models.Repo, models.PullRequest, models.User) {
	baseRepo, headRepo, pull, user := c.GetAllCapturedArguments()
	return baseRepo[len(baseRepo)-1], headRepo[len(headRepo)-1], pull[len(pull)-1], user[len(user)-1]
}

func (c *MockCommandRunner_RunAutoplanCommand_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.Repo, _param2 []models.PullRequest, _param3 []models.User) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]models.Repo, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(models.Repo)
			}
		}
		if len(_params) > 1 {
			_param1 = make([]models.Repo, len(c.methodInvocations))
			for u, param := range _params[1] {
				_param1[u] = param.(models.Repo)
			}
		}
		if len(_params) > 2 {
			_param2 = make([]models.PullRequest, len(c.methodInvocations))
			for u, param := range _params[2] {
				_param2[u] = param.(models.PullRequest)
			}
		}
		if len(_params) > 3 {
			_param3 = make([]models.User, len(c.methodInvocations))
			for u, param := range _params[3] {
				_param3[u] = param.(models.User)
			}
		}
	}
	return
}

func (verifier *VerifierMockCommandRunner) RunCommentCommand(baseRepo models.Repo, maybeHeadRepo *models.Repo, maybePull *models.PullRequest, user models.User, pullNum int, cmd *events.CommentCommand) *MockCommandRunner_RunCommentCommand_OngoingVerification {
	_params := []pegomock.Param{baseRepo, maybeHeadRepo, maybePull, user, pullNum, cmd}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RunCommentCommand", _params, verifier.timeout)
	return &MockCommandRunner_RunCommentCommand_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommandRunner_RunCommentCommand_OngoingVerification struct {
	mock              *MockCommandRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommandRunner_RunCommentCommand_OngoingVerification) GetCapturedArguments() (models.Repo, *models.Repo, *models.PullRequest, models.User, int, *events.CommentCommand) {
	baseRepo, maybeHeadRepo, maybePull, user, pullNum, cmd := c.GetAllCapturedArguments()
	return baseRepo[len(baseRepo)-1], maybeHeadRepo[len(maybeHeadRepo)-1], maybePull[len(maybePull)-1], user[len(user)-1], pullNum[len(pullNum)-1], cmd[len(cmd)-1]
}

func (c *MockCommandRunner_RunCommentCommand_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []*models.Repo, _param2 []*models.PullRequest, _param3 []models.User, _param4 []int, _param5 []*events.CommentCommand) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]models.Repo, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(models.Repo)
			}
		}
		if len(_params) > 1 {
			_param1 = make([]*models.Repo, len(c.methodInvocations))
			for u, param := range _params[1] {
				_param1[u] = param.(*models.Repo)
			}
		}
		if len(_params) > 2 {
			_param2 = make([]*models.PullRequest, len(c.methodInvocations))
			for u, param := range _params[2] {
				_param2[u] = param.(*models.PullRequest)
			}
		}
		if len(_params) > 3 {
			_param3 = make([]models.User, len(c.methodInvocations))
			for u, param := range _params[3] {
				_param3[u] = param.(models.User)
			}
		}
		if len(_params) > 4 {
			_param4 = make([]int, len(c.methodInvocations))
			for u, param := range _params[4] {
				_param4[u] = param.(int)
			}
		}
		if len(_params) > 5 {
			_param5 = make([]*events.CommentCommand, len(c.methodInvocations))
			for u, param := range _params[5] {
				_param5[u] = param.(*events.CommentCommand)
			}
		}
	}
	return
}
