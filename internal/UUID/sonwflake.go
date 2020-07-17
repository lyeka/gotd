package UUID

import "github.com/bwmarrin/snowflake"

type IDGenerator struct {
	SnowNode *snowflake.Node
}

func NewNode() (*IDGenerator, error) {
	node, err := snowflake.NewNode(1) // todo 现在默认只有一个节点吗后续需要扩展
	if err != nil {
		return nil, err
	}

	return &IDGenerator{SnowNode: node}, nil
}

func (IDG *IDGenerator) Generate() snowflake.ID {
	return IDG.SnowNode.Generate()
}
