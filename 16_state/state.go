package state

import "fmt"

// Machine 状态机
type Machine struct {
	state IState
}

// SetState 更新状态
func (m *Machine) SetState(state IState) {
	m.state = state
}

// GetStateName 获取当前状态
func (m *Machine) GetStateName() string {
	return m.state.GetName()
}

func (m *Machine) Approval() {
	m.state.Approval(m)
}

func (m *Machine) Reject() {
	m.state.Reject(m)
}

// IState 状态
type IState interface {
	// 审批通过
	Approval(m *Machine)
	// 驳回
	Reject(m *Machine)
	// 获取当前状态名称
	GetName() string
}

// leaderApproveState 直属领导审批
type leaderApproveState struct{}

// Approval 获取状态名字
func (leaderApproveState) Approval(m *Machine) {
	fmt.Println("leader 审批成功")
	m.SetState(GetFinanceApproveState())
}

// GetName 获取状态名字
func (leaderApproveState) GetName() string {
	return "LeaderApproveState"
}

// Reject 获取状态名字
func (leaderApproveState) Reject(m *Machine) {}

func GetLeaderApproveState() IState {
	return &leaderApproveState{}
}

// financeApproveState 财务审批
type financeApproveState struct{}

// Approval 审批通过
func (f financeApproveState) Approval(m *Machine) {
	fmt.Println("财务审批成功")
	fmt.Println("出发打款操作")
}

// 拒绝
func (f financeApproveState) Reject(m *Machine) {
	m.SetState(GetLeaderApproveState())
}

// GetName 获取名字
func (f financeApproveState) GetName() string {
	return "FinanceApproveState"
}

// GetFinanceApproveState GetFinanceApproveState
func GetFinanceApproveState() IState {
	return &financeApproveState{}
}
