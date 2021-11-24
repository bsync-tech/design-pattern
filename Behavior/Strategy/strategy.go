package strategy

type IStrategy interface {
	Process() string
}
type StrategyA struct {
}

func (s *StrategyA) Process() string {
	return "process with StrategyA"
}

type StrategyB struct {
}

func (s *StrategyB) Process() string {
	return "process with StrategyB"
}

type Context struct {
	strategy IStrategy
}

func (cx *Context) SetSrategy(strategy IStrategy) {
	cx.strategy = strategy
}
func (cx *Context) GetSrategy() IStrategy {
	return cx.strategy
}

func (cx *Context) Process() string {
	return cx.strategy.Process()
}
