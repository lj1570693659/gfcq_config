package cmd

var (
// Main is the main command.
//Main = gcmd.Command{
//	Name:  "gfcq_config",
//	Usage: "main",
//	Brief: "start grpc server of product config and inspirit config",
//	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
//		c := grpcx.Server.NewConfig()
//		c.Options = append(c.Options, []grpc.ServerOption{
//			grpcx.Server.ChainUnary(
//				grpcx.Server.UnaryValidate,
//				grpcx.Server.UnaryTracing,
//				grpcx.Server.UnaryError,
//				grpcx.Server.UnaryRecover,
//			)}...,
//		)
//		s := grpcx.Server.New(c)
//		// 项目配置
//		product.LevelAssessRegister(s)
//		product.LevelConfirmRegister(s)
//		product.ModeRegister(s)
//		product.TypeRegister(s)
//		product.ModeStageRegister(s)
//		product.RolesRegister(s)
//		// 激励配置
//		inspirit.BudgetAccessRegister(s)
//		inspirit.CrewDutyIndexRegister(s)
//		inspirit.CrewHoursIndexRegister(s)
//		inspirit.CrewManageIndexRegister(s)
//		inspirit.CrewOvertimeRuleRegister(s)
//		inspirit.CrewSolveRuleRegister(s)
//		inspirit.CrewKpiRuleRegister(s)
//		inspirit.ProductStageRadioRegister(s)
//		s.Run()
//		return nil
//	},
//}
)
