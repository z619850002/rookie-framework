package skeleton

import "hub000.xindong.com/rookie/rookie-framework/logic"

//Skeleton is the system that control the process in modules.
type Skeleton struct {
	LogicBlock 	logic.LogicBlock
}

//NewSkeleton creates a new skeleton, the block in it is the base logic block.
func NewBaseSkeleton() Skeleton{
	return Skeleton{LogicBlock:logic.NewBaseLogicBlock()}
}

//WrapLog add log function to the process wrap in this module.
func (h * Skeleton) WrapLog(){
	h.LogicBlock = logic.NewLogWrapper(h.LogicBlock)
}

//WrapMemory add memory statistic function to the process wrap in this module.
func (h * Skeleton) WrapMemory(){
	h.LogicBlock = logic.NewMemoryWrapper(h.LogicBlock)
}