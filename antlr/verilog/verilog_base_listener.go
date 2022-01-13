// Code generated from Verilog.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Verilog

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVerilogListener is a complete listener for a parse tree produced by VerilogParser.
type BaseVerilogListener struct{}

var _ VerilogListener = &BaseVerilogListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVerilogListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVerilogListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVerilogListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVerilogListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterConfig_declaration is called when production config_declaration is entered.
func (s *BaseVerilogListener) EnterConfig_declaration(ctx *Config_declarationContext) {}

// ExitConfig_declaration is called when production config_declaration is exited.
func (s *BaseVerilogListener) ExitConfig_declaration(ctx *Config_declarationContext) {}

// EnterDesign_statement is called when production design_statement is entered.
func (s *BaseVerilogListener) EnterDesign_statement(ctx *Design_statementContext) {}

// ExitDesign_statement is called when production design_statement is exited.
func (s *BaseVerilogListener) ExitDesign_statement(ctx *Design_statementContext) {}

// EnterConfig_rule_statement is called when production config_rule_statement is entered.
func (s *BaseVerilogListener) EnterConfig_rule_statement(ctx *Config_rule_statementContext) {}

// ExitConfig_rule_statement is called when production config_rule_statement is exited.
func (s *BaseVerilogListener) ExitConfig_rule_statement(ctx *Config_rule_statementContext) {}

// EnterDefault_clause is called when production default_clause is entered.
func (s *BaseVerilogListener) EnterDefault_clause(ctx *Default_clauseContext) {}

// ExitDefault_clause is called when production default_clause is exited.
func (s *BaseVerilogListener) ExitDefault_clause(ctx *Default_clauseContext) {}

// EnterInst_clause is called when production inst_clause is entered.
func (s *BaseVerilogListener) EnterInst_clause(ctx *Inst_clauseContext) {}

// ExitInst_clause is called when production inst_clause is exited.
func (s *BaseVerilogListener) ExitInst_clause(ctx *Inst_clauseContext) {}

// EnterInst_name is called when production inst_name is entered.
func (s *BaseVerilogListener) EnterInst_name(ctx *Inst_nameContext) {}

// ExitInst_name is called when production inst_name is exited.
func (s *BaseVerilogListener) ExitInst_name(ctx *Inst_nameContext) {}

// EnterLiblist_clause is called when production liblist_clause is entered.
func (s *BaseVerilogListener) EnterLiblist_clause(ctx *Liblist_clauseContext) {}

// ExitLiblist_clause is called when production liblist_clause is exited.
func (s *BaseVerilogListener) ExitLiblist_clause(ctx *Liblist_clauseContext) {}

// EnterCell_clause is called when production cell_clause is entered.
func (s *BaseVerilogListener) EnterCell_clause(ctx *Cell_clauseContext) {}

// ExitCell_clause is called when production cell_clause is exited.
func (s *BaseVerilogListener) ExitCell_clause(ctx *Cell_clauseContext) {}

// EnterUse_clause is called when production use_clause is entered.
func (s *BaseVerilogListener) EnterUse_clause(ctx *Use_clauseContext) {}

// ExitUse_clause is called when production use_clause is exited.
func (s *BaseVerilogListener) ExitUse_clause(ctx *Use_clauseContext) {}

// EnterSource_text is called when production source_text is entered.
func (s *BaseVerilogListener) EnterSource_text(ctx *Source_textContext) {}

// ExitSource_text is called when production source_text is exited.
func (s *BaseVerilogListener) ExitSource_text(ctx *Source_textContext) {}

// EnterDescription is called when production description is entered.
func (s *BaseVerilogListener) EnterDescription(ctx *DescriptionContext) {}

// ExitDescription is called when production description is exited.
func (s *BaseVerilogListener) ExitDescription(ctx *DescriptionContext) {}

// EnterModule_declaration is called when production module_declaration is entered.
func (s *BaseVerilogListener) EnterModule_declaration(ctx *Module_declarationContext) {}

// ExitModule_declaration is called when production module_declaration is exited.
func (s *BaseVerilogListener) ExitModule_declaration(ctx *Module_declarationContext) {}

// EnterModule_keyword is called when production module_keyword is entered.
func (s *BaseVerilogListener) EnterModule_keyword(ctx *Module_keywordContext) {}

// ExitModule_keyword is called when production module_keyword is exited.
func (s *BaseVerilogListener) ExitModule_keyword(ctx *Module_keywordContext) {}

// EnterModule_parameter_port_list is called when production module_parameter_port_list is entered.
func (s *BaseVerilogListener) EnterModule_parameter_port_list(ctx *Module_parameter_port_listContext) {
}

// ExitModule_parameter_port_list is called when production module_parameter_port_list is exited.
func (s *BaseVerilogListener) ExitModule_parameter_port_list(ctx *Module_parameter_port_listContext) {
}

// EnterList_of_ports is called when production list_of_ports is entered.
func (s *BaseVerilogListener) EnterList_of_ports(ctx *List_of_portsContext) {}

// ExitList_of_ports is called when production list_of_ports is exited.
func (s *BaseVerilogListener) ExitList_of_ports(ctx *List_of_portsContext) {}

// EnterList_of_port_declarations is called when production list_of_port_declarations is entered.
func (s *BaseVerilogListener) EnterList_of_port_declarations(ctx *List_of_port_declarationsContext) {}

// ExitList_of_port_declarations is called when production list_of_port_declarations is exited.
func (s *BaseVerilogListener) ExitList_of_port_declarations(ctx *List_of_port_declarationsContext) {}

// EnterPort is called when production port is entered.
func (s *BaseVerilogListener) EnterPort(ctx *PortContext) {}

// ExitPort is called when production port is exited.
func (s *BaseVerilogListener) ExitPort(ctx *PortContext) {}

// EnterPort_expression is called when production port_expression is entered.
func (s *BaseVerilogListener) EnterPort_expression(ctx *Port_expressionContext) {}

// ExitPort_expression is called when production port_expression is exited.
func (s *BaseVerilogListener) ExitPort_expression(ctx *Port_expressionContext) {}

// EnterPort_reference is called when production port_reference is entered.
func (s *BaseVerilogListener) EnterPort_reference(ctx *Port_referenceContext) {}

// ExitPort_reference is called when production port_reference is exited.
func (s *BaseVerilogListener) ExitPort_reference(ctx *Port_referenceContext) {}

// EnterPort_declaration is called when production port_declaration is entered.
func (s *BaseVerilogListener) EnterPort_declaration(ctx *Port_declarationContext) {}

// ExitPort_declaration is called when production port_declaration is exited.
func (s *BaseVerilogListener) ExitPort_declaration(ctx *Port_declarationContext) {}

// EnterModule_item is called when production module_item is entered.
func (s *BaseVerilogListener) EnterModule_item(ctx *Module_itemContext) {}

// ExitModule_item is called when production module_item is exited.
func (s *BaseVerilogListener) ExitModule_item(ctx *Module_itemContext) {}

// EnterModule_or_generate_item is called when production module_or_generate_item is entered.
func (s *BaseVerilogListener) EnterModule_or_generate_item(ctx *Module_or_generate_itemContext) {}

// ExitModule_or_generate_item is called when production module_or_generate_item is exited.
func (s *BaseVerilogListener) ExitModule_or_generate_item(ctx *Module_or_generate_itemContext) {}

// EnterNon_port_module_item is called when production non_port_module_item is entered.
func (s *BaseVerilogListener) EnterNon_port_module_item(ctx *Non_port_module_itemContext) {}

// ExitNon_port_module_item is called when production non_port_module_item is exited.
func (s *BaseVerilogListener) ExitNon_port_module_item(ctx *Non_port_module_itemContext) {}

// EnterModule_or_generate_item_declaration is called when production module_or_generate_item_declaration is entered.
func (s *BaseVerilogListener) EnterModule_or_generate_item_declaration(ctx *Module_or_generate_item_declarationContext) {
}

// ExitModule_or_generate_item_declaration is called when production module_or_generate_item_declaration is exited.
func (s *BaseVerilogListener) ExitModule_or_generate_item_declaration(ctx *Module_or_generate_item_declarationContext) {
}

// EnterParameter_override is called when production parameter_override is entered.
func (s *BaseVerilogListener) EnterParameter_override(ctx *Parameter_overrideContext) {}

// ExitParameter_override is called when production parameter_override is exited.
func (s *BaseVerilogListener) ExitParameter_override(ctx *Parameter_overrideContext) {}

// EnterLocal_parameter_declaration is called when production local_parameter_declaration is entered.
func (s *BaseVerilogListener) EnterLocal_parameter_declaration(ctx *Local_parameter_declarationContext) {
}

// ExitLocal_parameter_declaration is called when production local_parameter_declaration is exited.
func (s *BaseVerilogListener) ExitLocal_parameter_declaration(ctx *Local_parameter_declarationContext) {
}

// EnterParameter_declaration is called when production parameter_declaration is entered.
func (s *BaseVerilogListener) EnterParameter_declaration(ctx *Parameter_declarationContext) {}

// ExitParameter_declaration is called when production parameter_declaration is exited.
func (s *BaseVerilogListener) ExitParameter_declaration(ctx *Parameter_declarationContext) {}

// EnterParameter_declaration_ is called when production parameter_declaration_ is entered.
func (s *BaseVerilogListener) EnterParameter_declaration_(ctx *Parameter_declaration_Context) {}

// ExitParameter_declaration_ is called when production parameter_declaration_ is exited.
func (s *BaseVerilogListener) ExitParameter_declaration_(ctx *Parameter_declaration_Context) {}

// EnterSpecparam_declaration is called when production specparam_declaration is entered.
func (s *BaseVerilogListener) EnterSpecparam_declaration(ctx *Specparam_declarationContext) {}

// ExitSpecparam_declaration is called when production specparam_declaration is exited.
func (s *BaseVerilogListener) ExitSpecparam_declaration(ctx *Specparam_declarationContext) {}

// EnterInout_declaration is called when production inout_declaration is entered.
func (s *BaseVerilogListener) EnterInout_declaration(ctx *Inout_declarationContext) {}

// ExitInout_declaration is called when production inout_declaration is exited.
func (s *BaseVerilogListener) ExitInout_declaration(ctx *Inout_declarationContext) {}

// EnterInput_declaration is called when production input_declaration is entered.
func (s *BaseVerilogListener) EnterInput_declaration(ctx *Input_declarationContext) {}

// ExitInput_declaration is called when production input_declaration is exited.
func (s *BaseVerilogListener) ExitInput_declaration(ctx *Input_declarationContext) {}

// EnterOutput_declaration is called when production output_declaration is entered.
func (s *BaseVerilogListener) EnterOutput_declaration(ctx *Output_declarationContext) {}

// ExitOutput_declaration is called when production output_declaration is exited.
func (s *BaseVerilogListener) ExitOutput_declaration(ctx *Output_declarationContext) {}

// EnterEvent_declaration is called when production event_declaration is entered.
func (s *BaseVerilogListener) EnterEvent_declaration(ctx *Event_declarationContext) {}

// ExitEvent_declaration is called when production event_declaration is exited.
func (s *BaseVerilogListener) ExitEvent_declaration(ctx *Event_declarationContext) {}

// EnterGenvar_declaration is called when production genvar_declaration is entered.
func (s *BaseVerilogListener) EnterGenvar_declaration(ctx *Genvar_declarationContext) {}

// ExitGenvar_declaration is called when production genvar_declaration is exited.
func (s *BaseVerilogListener) ExitGenvar_declaration(ctx *Genvar_declarationContext) {}

// EnterInteger_declaration is called when production integer_declaration is entered.
func (s *BaseVerilogListener) EnterInteger_declaration(ctx *Integer_declarationContext) {}

// ExitInteger_declaration is called when production integer_declaration is exited.
func (s *BaseVerilogListener) ExitInteger_declaration(ctx *Integer_declarationContext) {}

// EnterTime_declaration is called when production time_declaration is entered.
func (s *BaseVerilogListener) EnterTime_declaration(ctx *Time_declarationContext) {}

// ExitTime_declaration is called when production time_declaration is exited.
func (s *BaseVerilogListener) ExitTime_declaration(ctx *Time_declarationContext) {}

// EnterReal_declaration is called when production real_declaration is entered.
func (s *BaseVerilogListener) EnterReal_declaration(ctx *Real_declarationContext) {}

// ExitReal_declaration is called when production real_declaration is exited.
func (s *BaseVerilogListener) ExitReal_declaration(ctx *Real_declarationContext) {}

// EnterRealtime_declaration is called when production realtime_declaration is entered.
func (s *BaseVerilogListener) EnterRealtime_declaration(ctx *Realtime_declarationContext) {}

// ExitRealtime_declaration is called when production realtime_declaration is exited.
func (s *BaseVerilogListener) ExitRealtime_declaration(ctx *Realtime_declarationContext) {}

// EnterReg_declaration is called when production reg_declaration is entered.
func (s *BaseVerilogListener) EnterReg_declaration(ctx *Reg_declarationContext) {}

// ExitReg_declaration is called when production reg_declaration is exited.
func (s *BaseVerilogListener) ExitReg_declaration(ctx *Reg_declarationContext) {}

// EnterNet_declaration is called when production net_declaration is entered.
func (s *BaseVerilogListener) EnterNet_declaration(ctx *Net_declarationContext) {}

// ExitNet_declaration is called when production net_declaration is exited.
func (s *BaseVerilogListener) ExitNet_declaration(ctx *Net_declarationContext) {}

// EnterNet_type is called when production net_type is entered.
func (s *BaseVerilogListener) EnterNet_type(ctx *Net_typeContext) {}

// ExitNet_type is called when production net_type is exited.
func (s *BaseVerilogListener) ExitNet_type(ctx *Net_typeContext) {}

// EnterOutput_variable_type is called when production output_variable_type is entered.
func (s *BaseVerilogListener) EnterOutput_variable_type(ctx *Output_variable_typeContext) {}

// ExitOutput_variable_type is called when production output_variable_type is exited.
func (s *BaseVerilogListener) ExitOutput_variable_type(ctx *Output_variable_typeContext) {}

// EnterReal_type is called when production real_type is entered.
func (s *BaseVerilogListener) EnterReal_type(ctx *Real_typeContext) {}

// ExitReal_type is called when production real_type is exited.
func (s *BaseVerilogListener) ExitReal_type(ctx *Real_typeContext) {}

// EnterVariable_type is called when production variable_type is entered.
func (s *BaseVerilogListener) EnterVariable_type(ctx *Variable_typeContext) {}

// ExitVariable_type is called when production variable_type is exited.
func (s *BaseVerilogListener) ExitVariable_type(ctx *Variable_typeContext) {}

// EnterDrive_strength is called when production drive_strength is entered.
func (s *BaseVerilogListener) EnterDrive_strength(ctx *Drive_strengthContext) {}

// ExitDrive_strength is called when production drive_strength is exited.
func (s *BaseVerilogListener) ExitDrive_strength(ctx *Drive_strengthContext) {}

// EnterStrength0 is called when production strength0 is entered.
func (s *BaseVerilogListener) EnterStrength0(ctx *Strength0Context) {}

// ExitStrength0 is called when production strength0 is exited.
func (s *BaseVerilogListener) ExitStrength0(ctx *Strength0Context) {}

// EnterStrength1 is called when production strength1 is entered.
func (s *BaseVerilogListener) EnterStrength1(ctx *Strength1Context) {}

// ExitStrength1 is called when production strength1 is exited.
func (s *BaseVerilogListener) ExitStrength1(ctx *Strength1Context) {}

// EnterCharge_strength is called when production charge_strength is entered.
func (s *BaseVerilogListener) EnterCharge_strength(ctx *Charge_strengthContext) {}

// ExitCharge_strength is called when production charge_strength is exited.
func (s *BaseVerilogListener) ExitCharge_strength(ctx *Charge_strengthContext) {}

// EnterDelay3 is called when production delay3 is entered.
func (s *BaseVerilogListener) EnterDelay3(ctx *Delay3Context) {}

// ExitDelay3 is called when production delay3 is exited.
func (s *BaseVerilogListener) ExitDelay3(ctx *Delay3Context) {}

// EnterDelay2 is called when production delay2 is entered.
func (s *BaseVerilogListener) EnterDelay2(ctx *Delay2Context) {}

// ExitDelay2 is called when production delay2 is exited.
func (s *BaseVerilogListener) ExitDelay2(ctx *Delay2Context) {}

// EnterDelay_value is called when production delay_value is entered.
func (s *BaseVerilogListener) EnterDelay_value(ctx *Delay_valueContext) {}

// ExitDelay_value is called when production delay_value is exited.
func (s *BaseVerilogListener) ExitDelay_value(ctx *Delay_valueContext) {}

// EnterList_of_event_identifiers is called when production list_of_event_identifiers is entered.
func (s *BaseVerilogListener) EnterList_of_event_identifiers(ctx *List_of_event_identifiersContext) {}

// ExitList_of_event_identifiers is called when production list_of_event_identifiers is exited.
func (s *BaseVerilogListener) ExitList_of_event_identifiers(ctx *List_of_event_identifiersContext) {}

// EnterList_of_net_identifiers is called when production list_of_net_identifiers is entered.
func (s *BaseVerilogListener) EnterList_of_net_identifiers(ctx *List_of_net_identifiersContext) {}

// ExitList_of_net_identifiers is called when production list_of_net_identifiers is exited.
func (s *BaseVerilogListener) ExitList_of_net_identifiers(ctx *List_of_net_identifiersContext) {}

// EnterList_of_genvar_identifiers is called when production list_of_genvar_identifiers is entered.
func (s *BaseVerilogListener) EnterList_of_genvar_identifiers(ctx *List_of_genvar_identifiersContext) {
}

// ExitList_of_genvar_identifiers is called when production list_of_genvar_identifiers is exited.
func (s *BaseVerilogListener) ExitList_of_genvar_identifiers(ctx *List_of_genvar_identifiersContext) {
}

// EnterList_of_port_identifiers is called when production list_of_port_identifiers is entered.
func (s *BaseVerilogListener) EnterList_of_port_identifiers(ctx *List_of_port_identifiersContext) {}

// ExitList_of_port_identifiers is called when production list_of_port_identifiers is exited.
func (s *BaseVerilogListener) ExitList_of_port_identifiers(ctx *List_of_port_identifiersContext) {}

// EnterList_of_net_decl_assignments is called when production list_of_net_decl_assignments is entered.
func (s *BaseVerilogListener) EnterList_of_net_decl_assignments(ctx *List_of_net_decl_assignmentsContext) {
}

// ExitList_of_net_decl_assignments is called when production list_of_net_decl_assignments is exited.
func (s *BaseVerilogListener) ExitList_of_net_decl_assignments(ctx *List_of_net_decl_assignmentsContext) {
}

// EnterList_of_param_assignments is called when production list_of_param_assignments is entered.
func (s *BaseVerilogListener) EnterList_of_param_assignments(ctx *List_of_param_assignmentsContext) {}

// ExitList_of_param_assignments is called when production list_of_param_assignments is exited.
func (s *BaseVerilogListener) ExitList_of_param_assignments(ctx *List_of_param_assignmentsContext) {}

// EnterList_of_specparam_assignments is called when production list_of_specparam_assignments is entered.
func (s *BaseVerilogListener) EnterList_of_specparam_assignments(ctx *List_of_specparam_assignmentsContext) {
}

// ExitList_of_specparam_assignments is called when production list_of_specparam_assignments is exited.
func (s *BaseVerilogListener) ExitList_of_specparam_assignments(ctx *List_of_specparam_assignmentsContext) {
}

// EnterList_of_real_identifiers is called when production list_of_real_identifiers is entered.
func (s *BaseVerilogListener) EnterList_of_real_identifiers(ctx *List_of_real_identifiersContext) {}

// ExitList_of_real_identifiers is called when production list_of_real_identifiers is exited.
func (s *BaseVerilogListener) ExitList_of_real_identifiers(ctx *List_of_real_identifiersContext) {}

// EnterList_of_variable_identifiers is called when production list_of_variable_identifiers is entered.
func (s *BaseVerilogListener) EnterList_of_variable_identifiers(ctx *List_of_variable_identifiersContext) {
}

// ExitList_of_variable_identifiers is called when production list_of_variable_identifiers is exited.
func (s *BaseVerilogListener) ExitList_of_variable_identifiers(ctx *List_of_variable_identifiersContext) {
}

// EnterList_of_variable_port_identifiers is called when production list_of_variable_port_identifiers is entered.
func (s *BaseVerilogListener) EnterList_of_variable_port_identifiers(ctx *List_of_variable_port_identifiersContext) {
}

// ExitList_of_variable_port_identifiers is called when production list_of_variable_port_identifiers is exited.
func (s *BaseVerilogListener) ExitList_of_variable_port_identifiers(ctx *List_of_variable_port_identifiersContext) {
}

// EnterNet_decl_assignment is called when production net_decl_assignment is entered.
func (s *BaseVerilogListener) EnterNet_decl_assignment(ctx *Net_decl_assignmentContext) {}

// ExitNet_decl_assignment is called when production net_decl_assignment is exited.
func (s *BaseVerilogListener) ExitNet_decl_assignment(ctx *Net_decl_assignmentContext) {}

// EnterParam_assignment is called when production param_assignment is entered.
func (s *BaseVerilogListener) EnterParam_assignment(ctx *Param_assignmentContext) {}

// ExitParam_assignment is called when production param_assignment is exited.
func (s *BaseVerilogListener) ExitParam_assignment(ctx *Param_assignmentContext) {}

// EnterSpecparam_assignment is called when production specparam_assignment is entered.
func (s *BaseVerilogListener) EnterSpecparam_assignment(ctx *Specparam_assignmentContext) {}

// ExitSpecparam_assignment is called when production specparam_assignment is exited.
func (s *BaseVerilogListener) ExitSpecparam_assignment(ctx *Specparam_assignmentContext) {}

// EnterPulse_control_specparam is called when production pulse_control_specparam is entered.
func (s *BaseVerilogListener) EnterPulse_control_specparam(ctx *Pulse_control_specparamContext) {}

// ExitPulse_control_specparam is called when production pulse_control_specparam is exited.
func (s *BaseVerilogListener) ExitPulse_control_specparam(ctx *Pulse_control_specparamContext) {}

// EnterError_limit_value is called when production error_limit_value is entered.
func (s *BaseVerilogListener) EnterError_limit_value(ctx *Error_limit_valueContext) {}

// ExitError_limit_value is called when production error_limit_value is exited.
func (s *BaseVerilogListener) ExitError_limit_value(ctx *Error_limit_valueContext) {}

// EnterReject_limit_value is called when production reject_limit_value is entered.
func (s *BaseVerilogListener) EnterReject_limit_value(ctx *Reject_limit_valueContext) {}

// ExitReject_limit_value is called when production reject_limit_value is exited.
func (s *BaseVerilogListener) ExitReject_limit_value(ctx *Reject_limit_valueContext) {}

// EnterLimit_value is called when production limit_value is entered.
func (s *BaseVerilogListener) EnterLimit_value(ctx *Limit_valueContext) {}

// ExitLimit_value is called when production limit_value is exited.
func (s *BaseVerilogListener) ExitLimit_value(ctx *Limit_valueContext) {}

// EnterDimension is called when production dimension is entered.
func (s *BaseVerilogListener) EnterDimension(ctx *DimensionContext) {}

// ExitDimension is called when production dimension is exited.
func (s *BaseVerilogListener) ExitDimension(ctx *DimensionContext) {}

// EnterRange_ is called when production range_ is entered.
func (s *BaseVerilogListener) EnterRange_(ctx *Range_Context) {}

// ExitRange_ is called when production range_ is exited.
func (s *BaseVerilogListener) ExitRange_(ctx *Range_Context) {}

// EnterFunction_declaration is called when production function_declaration is entered.
func (s *BaseVerilogListener) EnterFunction_declaration(ctx *Function_declarationContext) {}

// ExitFunction_declaration is called when production function_declaration is exited.
func (s *BaseVerilogListener) ExitFunction_declaration(ctx *Function_declarationContext) {}

// EnterFunction_item_declaration is called when production function_item_declaration is entered.
func (s *BaseVerilogListener) EnterFunction_item_declaration(ctx *Function_item_declarationContext) {}

// ExitFunction_item_declaration is called when production function_item_declaration is exited.
func (s *BaseVerilogListener) ExitFunction_item_declaration(ctx *Function_item_declarationContext) {}

// EnterFunction_port_list is called when production function_port_list is entered.
func (s *BaseVerilogListener) EnterFunction_port_list(ctx *Function_port_listContext) {}

// ExitFunction_port_list is called when production function_port_list is exited.
func (s *BaseVerilogListener) ExitFunction_port_list(ctx *Function_port_listContext) {}

// EnterFunction_port is called when production function_port is entered.
func (s *BaseVerilogListener) EnterFunction_port(ctx *Function_portContext) {}

// ExitFunction_port is called when production function_port is exited.
func (s *BaseVerilogListener) ExitFunction_port(ctx *Function_portContext) {}

// EnterRange_or_type is called when production range_or_type is entered.
func (s *BaseVerilogListener) EnterRange_or_type(ctx *Range_or_typeContext) {}

// ExitRange_or_type is called when production range_or_type is exited.
func (s *BaseVerilogListener) ExitRange_or_type(ctx *Range_or_typeContext) {}

// EnterTask_declaration is called when production task_declaration is entered.
func (s *BaseVerilogListener) EnterTask_declaration(ctx *Task_declarationContext) {}

// ExitTask_declaration is called when production task_declaration is exited.
func (s *BaseVerilogListener) ExitTask_declaration(ctx *Task_declarationContext) {}

// EnterTask_item_declaration is called when production task_item_declaration is entered.
func (s *BaseVerilogListener) EnterTask_item_declaration(ctx *Task_item_declarationContext) {}

// ExitTask_item_declaration is called when production task_item_declaration is exited.
func (s *BaseVerilogListener) ExitTask_item_declaration(ctx *Task_item_declarationContext) {}

// EnterTask_port_list is called when production task_port_list is entered.
func (s *BaseVerilogListener) EnterTask_port_list(ctx *Task_port_listContext) {}

// ExitTask_port_list is called when production task_port_list is exited.
func (s *BaseVerilogListener) ExitTask_port_list(ctx *Task_port_listContext) {}

// EnterTask_port_item is called when production task_port_item is entered.
func (s *BaseVerilogListener) EnterTask_port_item(ctx *Task_port_itemContext) {}

// ExitTask_port_item is called when production task_port_item is exited.
func (s *BaseVerilogListener) ExitTask_port_item(ctx *Task_port_itemContext) {}

// EnterTf_decl_header is called when production tf_decl_header is entered.
func (s *BaseVerilogListener) EnterTf_decl_header(ctx *Tf_decl_headerContext) {}

// ExitTf_decl_header is called when production tf_decl_header is exited.
func (s *BaseVerilogListener) ExitTf_decl_header(ctx *Tf_decl_headerContext) {}

// EnterTf_declaration is called when production tf_declaration is entered.
func (s *BaseVerilogListener) EnterTf_declaration(ctx *Tf_declarationContext) {}

// ExitTf_declaration is called when production tf_declaration is exited.
func (s *BaseVerilogListener) ExitTf_declaration(ctx *Tf_declarationContext) {}

// EnterTask_port_type is called when production task_port_type is entered.
func (s *BaseVerilogListener) EnterTask_port_type(ctx *Task_port_typeContext) {}

// ExitTask_port_type is called when production task_port_type is exited.
func (s *BaseVerilogListener) ExitTask_port_type(ctx *Task_port_typeContext) {}

// EnterBlock_item_declaration is called when production block_item_declaration is entered.
func (s *BaseVerilogListener) EnterBlock_item_declaration(ctx *Block_item_declarationContext) {}

// ExitBlock_item_declaration is called when production block_item_declaration is exited.
func (s *BaseVerilogListener) ExitBlock_item_declaration(ctx *Block_item_declarationContext) {}

// EnterBlock_reg_declaration is called when production block_reg_declaration is entered.
func (s *BaseVerilogListener) EnterBlock_reg_declaration(ctx *Block_reg_declarationContext) {}

// ExitBlock_reg_declaration is called when production block_reg_declaration is exited.
func (s *BaseVerilogListener) ExitBlock_reg_declaration(ctx *Block_reg_declarationContext) {}

// EnterList_of_block_variable_identifiers is called when production list_of_block_variable_identifiers is entered.
func (s *BaseVerilogListener) EnterList_of_block_variable_identifiers(ctx *List_of_block_variable_identifiersContext) {
}

// ExitList_of_block_variable_identifiers is called when production list_of_block_variable_identifiers is exited.
func (s *BaseVerilogListener) ExitList_of_block_variable_identifiers(ctx *List_of_block_variable_identifiersContext) {
}

// EnterBlock_variable_type is called when production block_variable_type is entered.
func (s *BaseVerilogListener) EnterBlock_variable_type(ctx *Block_variable_typeContext) {}

// ExitBlock_variable_type is called when production block_variable_type is exited.
func (s *BaseVerilogListener) ExitBlock_variable_type(ctx *Block_variable_typeContext) {}

// EnterGate_instantiation is called when production gate_instantiation is entered.
func (s *BaseVerilogListener) EnterGate_instantiation(ctx *Gate_instantiationContext) {}

// ExitGate_instantiation is called when production gate_instantiation is exited.
func (s *BaseVerilogListener) ExitGate_instantiation(ctx *Gate_instantiationContext) {}

// EnterCmos_switch_instance is called when production cmos_switch_instance is entered.
func (s *BaseVerilogListener) EnterCmos_switch_instance(ctx *Cmos_switch_instanceContext) {}

// ExitCmos_switch_instance is called when production cmos_switch_instance is exited.
func (s *BaseVerilogListener) ExitCmos_switch_instance(ctx *Cmos_switch_instanceContext) {}

// EnterEnable_gate_instance is called when production enable_gate_instance is entered.
func (s *BaseVerilogListener) EnterEnable_gate_instance(ctx *Enable_gate_instanceContext) {}

// ExitEnable_gate_instance is called when production enable_gate_instance is exited.
func (s *BaseVerilogListener) ExitEnable_gate_instance(ctx *Enable_gate_instanceContext) {}

// EnterMos_switch_instance is called when production mos_switch_instance is entered.
func (s *BaseVerilogListener) EnterMos_switch_instance(ctx *Mos_switch_instanceContext) {}

// ExitMos_switch_instance is called when production mos_switch_instance is exited.
func (s *BaseVerilogListener) ExitMos_switch_instance(ctx *Mos_switch_instanceContext) {}

// EnterN_input_gate_instance is called when production n_input_gate_instance is entered.
func (s *BaseVerilogListener) EnterN_input_gate_instance(ctx *N_input_gate_instanceContext) {}

// ExitN_input_gate_instance is called when production n_input_gate_instance is exited.
func (s *BaseVerilogListener) ExitN_input_gate_instance(ctx *N_input_gate_instanceContext) {}

// EnterN_output_gate_instance is called when production n_output_gate_instance is entered.
func (s *BaseVerilogListener) EnterN_output_gate_instance(ctx *N_output_gate_instanceContext) {}

// ExitN_output_gate_instance is called when production n_output_gate_instance is exited.
func (s *BaseVerilogListener) ExitN_output_gate_instance(ctx *N_output_gate_instanceContext) {}

// EnterPass_switch_instance is called when production pass_switch_instance is entered.
func (s *BaseVerilogListener) EnterPass_switch_instance(ctx *Pass_switch_instanceContext) {}

// ExitPass_switch_instance is called when production pass_switch_instance is exited.
func (s *BaseVerilogListener) ExitPass_switch_instance(ctx *Pass_switch_instanceContext) {}

// EnterPass_enable_switch_instance is called when production pass_enable_switch_instance is entered.
func (s *BaseVerilogListener) EnterPass_enable_switch_instance(ctx *Pass_enable_switch_instanceContext) {
}

// ExitPass_enable_switch_instance is called when production pass_enable_switch_instance is exited.
func (s *BaseVerilogListener) ExitPass_enable_switch_instance(ctx *Pass_enable_switch_instanceContext) {
}

// EnterPull_gate_instance is called when production pull_gate_instance is entered.
func (s *BaseVerilogListener) EnterPull_gate_instance(ctx *Pull_gate_instanceContext) {}

// ExitPull_gate_instance is called when production pull_gate_instance is exited.
func (s *BaseVerilogListener) ExitPull_gate_instance(ctx *Pull_gate_instanceContext) {}

// EnterName_of_gate_instance is called when production name_of_gate_instance is entered.
func (s *BaseVerilogListener) EnterName_of_gate_instance(ctx *Name_of_gate_instanceContext) {}

// ExitName_of_gate_instance is called when production name_of_gate_instance is exited.
func (s *BaseVerilogListener) ExitName_of_gate_instance(ctx *Name_of_gate_instanceContext) {}

// EnterPulldown_strength is called when production pulldown_strength is entered.
func (s *BaseVerilogListener) EnterPulldown_strength(ctx *Pulldown_strengthContext) {}

// ExitPulldown_strength is called when production pulldown_strength is exited.
func (s *BaseVerilogListener) ExitPulldown_strength(ctx *Pulldown_strengthContext) {}

// EnterPullup_strength is called when production pullup_strength is entered.
func (s *BaseVerilogListener) EnterPullup_strength(ctx *Pullup_strengthContext) {}

// ExitPullup_strength is called when production pullup_strength is exited.
func (s *BaseVerilogListener) ExitPullup_strength(ctx *Pullup_strengthContext) {}

// EnterEnable_terminal is called when production enable_terminal is entered.
func (s *BaseVerilogListener) EnterEnable_terminal(ctx *Enable_terminalContext) {}

// ExitEnable_terminal is called when production enable_terminal is exited.
func (s *BaseVerilogListener) ExitEnable_terminal(ctx *Enable_terminalContext) {}

// EnterNcontrol_terminal is called when production ncontrol_terminal is entered.
func (s *BaseVerilogListener) EnterNcontrol_terminal(ctx *Ncontrol_terminalContext) {}

// ExitNcontrol_terminal is called when production ncontrol_terminal is exited.
func (s *BaseVerilogListener) ExitNcontrol_terminal(ctx *Ncontrol_terminalContext) {}

// EnterPcontrol_terminal is called when production pcontrol_terminal is entered.
func (s *BaseVerilogListener) EnterPcontrol_terminal(ctx *Pcontrol_terminalContext) {}

// ExitPcontrol_terminal is called when production pcontrol_terminal is exited.
func (s *BaseVerilogListener) ExitPcontrol_terminal(ctx *Pcontrol_terminalContext) {}

// EnterInput_terminal is called when production input_terminal is entered.
func (s *BaseVerilogListener) EnterInput_terminal(ctx *Input_terminalContext) {}

// ExitInput_terminal is called when production input_terminal is exited.
func (s *BaseVerilogListener) ExitInput_terminal(ctx *Input_terminalContext) {}

// EnterInout_terminal is called when production inout_terminal is entered.
func (s *BaseVerilogListener) EnterInout_terminal(ctx *Inout_terminalContext) {}

// ExitInout_terminal is called when production inout_terminal is exited.
func (s *BaseVerilogListener) ExitInout_terminal(ctx *Inout_terminalContext) {}

// EnterOutput_terminal is called when production output_terminal is entered.
func (s *BaseVerilogListener) EnterOutput_terminal(ctx *Output_terminalContext) {}

// ExitOutput_terminal is called when production output_terminal is exited.
func (s *BaseVerilogListener) ExitOutput_terminal(ctx *Output_terminalContext) {}

// EnterCmos_switchtype is called when production cmos_switchtype is entered.
func (s *BaseVerilogListener) EnterCmos_switchtype(ctx *Cmos_switchtypeContext) {}

// ExitCmos_switchtype is called when production cmos_switchtype is exited.
func (s *BaseVerilogListener) ExitCmos_switchtype(ctx *Cmos_switchtypeContext) {}

// EnterEnable_gatetype is called when production enable_gatetype is entered.
func (s *BaseVerilogListener) EnterEnable_gatetype(ctx *Enable_gatetypeContext) {}

// ExitEnable_gatetype is called when production enable_gatetype is exited.
func (s *BaseVerilogListener) ExitEnable_gatetype(ctx *Enable_gatetypeContext) {}

// EnterMos_switchtype is called when production mos_switchtype is entered.
func (s *BaseVerilogListener) EnterMos_switchtype(ctx *Mos_switchtypeContext) {}

// ExitMos_switchtype is called when production mos_switchtype is exited.
func (s *BaseVerilogListener) ExitMos_switchtype(ctx *Mos_switchtypeContext) {}

// EnterN_input_gatetype is called when production n_input_gatetype is entered.
func (s *BaseVerilogListener) EnterN_input_gatetype(ctx *N_input_gatetypeContext) {}

// ExitN_input_gatetype is called when production n_input_gatetype is exited.
func (s *BaseVerilogListener) ExitN_input_gatetype(ctx *N_input_gatetypeContext) {}

// EnterN_output_gatetype is called when production n_output_gatetype is entered.
func (s *BaseVerilogListener) EnterN_output_gatetype(ctx *N_output_gatetypeContext) {}

// ExitN_output_gatetype is called when production n_output_gatetype is exited.
func (s *BaseVerilogListener) ExitN_output_gatetype(ctx *N_output_gatetypeContext) {}

// EnterPass_en_switchtype is called when production pass_en_switchtype is entered.
func (s *BaseVerilogListener) EnterPass_en_switchtype(ctx *Pass_en_switchtypeContext) {}

// ExitPass_en_switchtype is called when production pass_en_switchtype is exited.
func (s *BaseVerilogListener) ExitPass_en_switchtype(ctx *Pass_en_switchtypeContext) {}

// EnterPass_switchtype is called when production pass_switchtype is entered.
func (s *BaseVerilogListener) EnterPass_switchtype(ctx *Pass_switchtypeContext) {}

// ExitPass_switchtype is called when production pass_switchtype is exited.
func (s *BaseVerilogListener) ExitPass_switchtype(ctx *Pass_switchtypeContext) {}

// EnterModule_instantiation is called when production module_instantiation is entered.
func (s *BaseVerilogListener) EnterModule_instantiation(ctx *Module_instantiationContext) {}

// ExitModule_instantiation is called when production module_instantiation is exited.
func (s *BaseVerilogListener) ExitModule_instantiation(ctx *Module_instantiationContext) {}

// EnterParameter_value_assignment is called when production parameter_value_assignment is entered.
func (s *BaseVerilogListener) EnterParameter_value_assignment(ctx *Parameter_value_assignmentContext) {
}

// ExitParameter_value_assignment is called when production parameter_value_assignment is exited.
func (s *BaseVerilogListener) ExitParameter_value_assignment(ctx *Parameter_value_assignmentContext) {
}

// EnterList_of_parameter_assignments is called when production list_of_parameter_assignments is entered.
func (s *BaseVerilogListener) EnterList_of_parameter_assignments(ctx *List_of_parameter_assignmentsContext) {
}

// ExitList_of_parameter_assignments is called when production list_of_parameter_assignments is exited.
func (s *BaseVerilogListener) ExitList_of_parameter_assignments(ctx *List_of_parameter_assignmentsContext) {
}

// EnterOrdered_parameter_assignment is called when production ordered_parameter_assignment is entered.
func (s *BaseVerilogListener) EnterOrdered_parameter_assignment(ctx *Ordered_parameter_assignmentContext) {
}

// ExitOrdered_parameter_assignment is called when production ordered_parameter_assignment is exited.
func (s *BaseVerilogListener) ExitOrdered_parameter_assignment(ctx *Ordered_parameter_assignmentContext) {
}

// EnterNamed_parameter_assignment is called when production named_parameter_assignment is entered.
func (s *BaseVerilogListener) EnterNamed_parameter_assignment(ctx *Named_parameter_assignmentContext) {
}

// ExitNamed_parameter_assignment is called when production named_parameter_assignment is exited.
func (s *BaseVerilogListener) ExitNamed_parameter_assignment(ctx *Named_parameter_assignmentContext) {
}

// EnterModule_instance is called when production module_instance is entered.
func (s *BaseVerilogListener) EnterModule_instance(ctx *Module_instanceContext) {}

// ExitModule_instance is called when production module_instance is exited.
func (s *BaseVerilogListener) ExitModule_instance(ctx *Module_instanceContext) {}

// EnterName_of_instance is called when production name_of_instance is entered.
func (s *BaseVerilogListener) EnterName_of_instance(ctx *Name_of_instanceContext) {}

// ExitName_of_instance is called when production name_of_instance is exited.
func (s *BaseVerilogListener) ExitName_of_instance(ctx *Name_of_instanceContext) {}

// EnterList_of_port_connections is called when production list_of_port_connections is entered.
func (s *BaseVerilogListener) EnterList_of_port_connections(ctx *List_of_port_connectionsContext) {}

// ExitList_of_port_connections is called when production list_of_port_connections is exited.
func (s *BaseVerilogListener) ExitList_of_port_connections(ctx *List_of_port_connectionsContext) {}

// EnterOrdered_port_connection is called when production ordered_port_connection is entered.
func (s *BaseVerilogListener) EnterOrdered_port_connection(ctx *Ordered_port_connectionContext) {}

// ExitOrdered_port_connection is called when production ordered_port_connection is exited.
func (s *BaseVerilogListener) ExitOrdered_port_connection(ctx *Ordered_port_connectionContext) {}

// EnterNamed_port_connection is called when production named_port_connection is entered.
func (s *BaseVerilogListener) EnterNamed_port_connection(ctx *Named_port_connectionContext) {}

// ExitNamed_port_connection is called when production named_port_connection is exited.
func (s *BaseVerilogListener) ExitNamed_port_connection(ctx *Named_port_connectionContext) {}

// EnterGenerated_instantiation is called when production generated_instantiation is entered.
func (s *BaseVerilogListener) EnterGenerated_instantiation(ctx *Generated_instantiationContext) {}

// ExitGenerated_instantiation is called when production generated_instantiation is exited.
func (s *BaseVerilogListener) ExitGenerated_instantiation(ctx *Generated_instantiationContext) {}

// EnterGenerate_item_or_null is called when production generate_item_or_null is entered.
func (s *BaseVerilogListener) EnterGenerate_item_or_null(ctx *Generate_item_or_nullContext) {}

// ExitGenerate_item_or_null is called when production generate_item_or_null is exited.
func (s *BaseVerilogListener) ExitGenerate_item_or_null(ctx *Generate_item_or_nullContext) {}

// EnterGenerate_item is called when production generate_item is entered.
func (s *BaseVerilogListener) EnterGenerate_item(ctx *Generate_itemContext) {}

// ExitGenerate_item is called when production generate_item is exited.
func (s *BaseVerilogListener) ExitGenerate_item(ctx *Generate_itemContext) {}

// EnterGenerate_conditional_statement is called when production generate_conditional_statement is entered.
func (s *BaseVerilogListener) EnterGenerate_conditional_statement(ctx *Generate_conditional_statementContext) {
}

// ExitGenerate_conditional_statement is called when production generate_conditional_statement is exited.
func (s *BaseVerilogListener) ExitGenerate_conditional_statement(ctx *Generate_conditional_statementContext) {
}

// EnterGenerate_case_statement is called when production generate_case_statement is entered.
func (s *BaseVerilogListener) EnterGenerate_case_statement(ctx *Generate_case_statementContext) {}

// ExitGenerate_case_statement is called when production generate_case_statement is exited.
func (s *BaseVerilogListener) ExitGenerate_case_statement(ctx *Generate_case_statementContext) {}

// EnterGenvar_case_item is called when production genvar_case_item is entered.
func (s *BaseVerilogListener) EnterGenvar_case_item(ctx *Genvar_case_itemContext) {}

// ExitGenvar_case_item is called when production genvar_case_item is exited.
func (s *BaseVerilogListener) ExitGenvar_case_item(ctx *Genvar_case_itemContext) {}

// EnterGenerate_loop_statement is called when production generate_loop_statement is entered.
func (s *BaseVerilogListener) EnterGenerate_loop_statement(ctx *Generate_loop_statementContext) {}

// ExitGenerate_loop_statement is called when production generate_loop_statement is exited.
func (s *BaseVerilogListener) ExitGenerate_loop_statement(ctx *Generate_loop_statementContext) {}

// EnterGenvar_assignment is called when production genvar_assignment is entered.
func (s *BaseVerilogListener) EnterGenvar_assignment(ctx *Genvar_assignmentContext) {}

// ExitGenvar_assignment is called when production genvar_assignment is exited.
func (s *BaseVerilogListener) ExitGenvar_assignment(ctx *Genvar_assignmentContext) {}

// EnterGenerate_block is called when production generate_block is entered.
func (s *BaseVerilogListener) EnterGenerate_block(ctx *Generate_blockContext) {}

// ExitGenerate_block is called when production generate_block is exited.
func (s *BaseVerilogListener) ExitGenerate_block(ctx *Generate_blockContext) {}

// EnterContinuous_assign is called when production continuous_assign is entered.
func (s *BaseVerilogListener) EnterContinuous_assign(ctx *Continuous_assignContext) {}

// ExitContinuous_assign is called when production continuous_assign is exited.
func (s *BaseVerilogListener) ExitContinuous_assign(ctx *Continuous_assignContext) {}

// EnterList_of_net_assignments is called when production list_of_net_assignments is entered.
func (s *BaseVerilogListener) EnterList_of_net_assignments(ctx *List_of_net_assignmentsContext) {}

// ExitList_of_net_assignments is called when production list_of_net_assignments is exited.
func (s *BaseVerilogListener) ExitList_of_net_assignments(ctx *List_of_net_assignmentsContext) {}

// EnterNet_assignment is called when production net_assignment is entered.
func (s *BaseVerilogListener) EnterNet_assignment(ctx *Net_assignmentContext) {}

// ExitNet_assignment is called when production net_assignment is exited.
func (s *BaseVerilogListener) ExitNet_assignment(ctx *Net_assignmentContext) {}

// EnterInitial_construct is called when production initial_construct is entered.
func (s *BaseVerilogListener) EnterInitial_construct(ctx *Initial_constructContext) {}

// ExitInitial_construct is called when production initial_construct is exited.
func (s *BaseVerilogListener) ExitInitial_construct(ctx *Initial_constructContext) {}

// EnterAlways_construct is called when production always_construct is entered.
func (s *BaseVerilogListener) EnterAlways_construct(ctx *Always_constructContext) {}

// ExitAlways_construct is called when production always_construct is exited.
func (s *BaseVerilogListener) ExitAlways_construct(ctx *Always_constructContext) {}

// EnterBlocking_assignment is called when production blocking_assignment is entered.
func (s *BaseVerilogListener) EnterBlocking_assignment(ctx *Blocking_assignmentContext) {}

// ExitBlocking_assignment is called when production blocking_assignment is exited.
func (s *BaseVerilogListener) ExitBlocking_assignment(ctx *Blocking_assignmentContext) {}

// EnterNonblocking_assignment is called when production nonblocking_assignment is entered.
func (s *BaseVerilogListener) EnterNonblocking_assignment(ctx *Nonblocking_assignmentContext) {}

// ExitNonblocking_assignment is called when production nonblocking_assignment is exited.
func (s *BaseVerilogListener) ExitNonblocking_assignment(ctx *Nonblocking_assignmentContext) {}

// EnterProcedural_continuous_assignments is called when production procedural_continuous_assignments is entered.
func (s *BaseVerilogListener) EnterProcedural_continuous_assignments(ctx *Procedural_continuous_assignmentsContext) {
}

// ExitProcedural_continuous_assignments is called when production procedural_continuous_assignments is exited.
func (s *BaseVerilogListener) ExitProcedural_continuous_assignments(ctx *Procedural_continuous_assignmentsContext) {
}

// EnterFunction_blocking_assignment is called when production function_blocking_assignment is entered.
func (s *BaseVerilogListener) EnterFunction_blocking_assignment(ctx *Function_blocking_assignmentContext) {
}

// ExitFunction_blocking_assignment is called when production function_blocking_assignment is exited.
func (s *BaseVerilogListener) ExitFunction_blocking_assignment(ctx *Function_blocking_assignmentContext) {
}

// EnterFunction_statement_or_null is called when production function_statement_or_null is entered.
func (s *BaseVerilogListener) EnterFunction_statement_or_null(ctx *Function_statement_or_nullContext) {
}

// ExitFunction_statement_or_null is called when production function_statement_or_null is exited.
func (s *BaseVerilogListener) ExitFunction_statement_or_null(ctx *Function_statement_or_nullContext) {
}

// EnterFunction_seq_block is called when production function_seq_block is entered.
func (s *BaseVerilogListener) EnterFunction_seq_block(ctx *Function_seq_blockContext) {}

// ExitFunction_seq_block is called when production function_seq_block is exited.
func (s *BaseVerilogListener) ExitFunction_seq_block(ctx *Function_seq_blockContext) {}

// EnterVariable_assignment is called when production variable_assignment is entered.
func (s *BaseVerilogListener) EnterVariable_assignment(ctx *Variable_assignmentContext) {}

// ExitVariable_assignment is called when production variable_assignment is exited.
func (s *BaseVerilogListener) ExitVariable_assignment(ctx *Variable_assignmentContext) {}

// EnterPar_block is called when production par_block is entered.
func (s *BaseVerilogListener) EnterPar_block(ctx *Par_blockContext) {}

// ExitPar_block is called when production par_block is exited.
func (s *BaseVerilogListener) ExitPar_block(ctx *Par_blockContext) {}

// EnterSeq_block is called when production seq_block is entered.
func (s *BaseVerilogListener) EnterSeq_block(ctx *Seq_blockContext) {}

// ExitSeq_block is called when production seq_block is exited.
func (s *BaseVerilogListener) ExitSeq_block(ctx *Seq_blockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseVerilogListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseVerilogListener) ExitStatement(ctx *StatementContext) {}

// EnterStatement_or_null is called when production statement_or_null is entered.
func (s *BaseVerilogListener) EnterStatement_or_null(ctx *Statement_or_nullContext) {}

// ExitStatement_or_null is called when production statement_or_null is exited.
func (s *BaseVerilogListener) ExitStatement_or_null(ctx *Statement_or_nullContext) {}

// EnterFunction_statement is called when production function_statement is entered.
func (s *BaseVerilogListener) EnterFunction_statement(ctx *Function_statementContext) {}

// ExitFunction_statement is called when production function_statement is exited.
func (s *BaseVerilogListener) ExitFunction_statement(ctx *Function_statementContext) {}

// EnterDelay_or_event_control is called when production delay_or_event_control is entered.
func (s *BaseVerilogListener) EnterDelay_or_event_control(ctx *Delay_or_event_controlContext) {}

// ExitDelay_or_event_control is called when production delay_or_event_control is exited.
func (s *BaseVerilogListener) ExitDelay_or_event_control(ctx *Delay_or_event_controlContext) {}

// EnterDelay_control is called when production delay_control is entered.
func (s *BaseVerilogListener) EnterDelay_control(ctx *Delay_controlContext) {}

// ExitDelay_control is called when production delay_control is exited.
func (s *BaseVerilogListener) ExitDelay_control(ctx *Delay_controlContext) {}

// EnterDisable_statement is called when production disable_statement is entered.
func (s *BaseVerilogListener) EnterDisable_statement(ctx *Disable_statementContext) {}

// ExitDisable_statement is called when production disable_statement is exited.
func (s *BaseVerilogListener) ExitDisable_statement(ctx *Disable_statementContext) {}

// EnterEvent_control is called when production event_control is entered.
func (s *BaseVerilogListener) EnterEvent_control(ctx *Event_controlContext) {}

// ExitEvent_control is called when production event_control is exited.
func (s *BaseVerilogListener) ExitEvent_control(ctx *Event_controlContext) {}

// EnterEvent_trigger is called when production event_trigger is entered.
func (s *BaseVerilogListener) EnterEvent_trigger(ctx *Event_triggerContext) {}

// ExitEvent_trigger is called when production event_trigger is exited.
func (s *BaseVerilogListener) ExitEvent_trigger(ctx *Event_triggerContext) {}

// EnterEvent_expression is called when production event_expression is entered.
func (s *BaseVerilogListener) EnterEvent_expression(ctx *Event_expressionContext) {}

// ExitEvent_expression is called when production event_expression is exited.
func (s *BaseVerilogListener) ExitEvent_expression(ctx *Event_expressionContext) {}

// EnterEvent_primary is called when production event_primary is entered.
func (s *BaseVerilogListener) EnterEvent_primary(ctx *Event_primaryContext) {}

// ExitEvent_primary is called when production event_primary is exited.
func (s *BaseVerilogListener) ExitEvent_primary(ctx *Event_primaryContext) {}

// EnterProcedural_timing_control_statement is called when production procedural_timing_control_statement is entered.
func (s *BaseVerilogListener) EnterProcedural_timing_control_statement(ctx *Procedural_timing_control_statementContext) {
}

// ExitProcedural_timing_control_statement is called when production procedural_timing_control_statement is exited.
func (s *BaseVerilogListener) ExitProcedural_timing_control_statement(ctx *Procedural_timing_control_statementContext) {
}

// EnterWait_statement is called when production wait_statement is entered.
func (s *BaseVerilogListener) EnterWait_statement(ctx *Wait_statementContext) {}

// ExitWait_statement is called when production wait_statement is exited.
func (s *BaseVerilogListener) ExitWait_statement(ctx *Wait_statementContext) {}

// EnterConditional_statement is called when production conditional_statement is entered.
func (s *BaseVerilogListener) EnterConditional_statement(ctx *Conditional_statementContext) {}

// ExitConditional_statement is called when production conditional_statement is exited.
func (s *BaseVerilogListener) ExitConditional_statement(ctx *Conditional_statementContext) {}

// EnterIf_else_if_statement is called when production if_else_if_statement is entered.
func (s *BaseVerilogListener) EnterIf_else_if_statement(ctx *If_else_if_statementContext) {}

// ExitIf_else_if_statement is called when production if_else_if_statement is exited.
func (s *BaseVerilogListener) ExitIf_else_if_statement(ctx *If_else_if_statementContext) {}

// EnterElse_statement is called when production else_statement is entered.
func (s *BaseVerilogListener) EnterElse_statement(ctx *Else_statementContext) {}

// ExitElse_statement is called when production else_statement is exited.
func (s *BaseVerilogListener) ExitElse_statement(ctx *Else_statementContext) {}

// EnterFunction_conditional_statement is called when production function_conditional_statement is entered.
func (s *BaseVerilogListener) EnterFunction_conditional_statement(ctx *Function_conditional_statementContext) {
}

// ExitFunction_conditional_statement is called when production function_conditional_statement is exited.
func (s *BaseVerilogListener) ExitFunction_conditional_statement(ctx *Function_conditional_statementContext) {
}

// EnterFunction_if_else_if_statement is called when production function_if_else_if_statement is entered.
func (s *BaseVerilogListener) EnterFunction_if_else_if_statement(ctx *Function_if_else_if_statementContext) {
}

// ExitFunction_if_else_if_statement is called when production function_if_else_if_statement is exited.
func (s *BaseVerilogListener) ExitFunction_if_else_if_statement(ctx *Function_if_else_if_statementContext) {
}

// EnterCase_statement is called when production case_statement is entered.
func (s *BaseVerilogListener) EnterCase_statement(ctx *Case_statementContext) {}

// ExitCase_statement is called when production case_statement is exited.
func (s *BaseVerilogListener) ExitCase_statement(ctx *Case_statementContext) {}

// EnterCase_item is called when production case_item is entered.
func (s *BaseVerilogListener) EnterCase_item(ctx *Case_itemContext) {}

// ExitCase_item is called when production case_item is exited.
func (s *BaseVerilogListener) ExitCase_item(ctx *Case_itemContext) {}

// EnterFunction_case_statement is called when production function_case_statement is entered.
func (s *BaseVerilogListener) EnterFunction_case_statement(ctx *Function_case_statementContext) {}

// ExitFunction_case_statement is called when production function_case_statement is exited.
func (s *BaseVerilogListener) ExitFunction_case_statement(ctx *Function_case_statementContext) {}

// EnterFunction_case_item is called when production function_case_item is entered.
func (s *BaseVerilogListener) EnterFunction_case_item(ctx *Function_case_itemContext) {}

// ExitFunction_case_item is called when production function_case_item is exited.
func (s *BaseVerilogListener) ExitFunction_case_item(ctx *Function_case_itemContext) {}

// EnterFunction_loop_statement is called when production function_loop_statement is entered.
func (s *BaseVerilogListener) EnterFunction_loop_statement(ctx *Function_loop_statementContext) {}

// ExitFunction_loop_statement is called when production function_loop_statement is exited.
func (s *BaseVerilogListener) ExitFunction_loop_statement(ctx *Function_loop_statementContext) {}

// EnterLoop_statement is called when production loop_statement is entered.
func (s *BaseVerilogListener) EnterLoop_statement(ctx *Loop_statementContext) {}

// ExitLoop_statement is called when production loop_statement is exited.
func (s *BaseVerilogListener) ExitLoop_statement(ctx *Loop_statementContext) {}

// EnterSystem_task_enable is called when production system_task_enable is entered.
func (s *BaseVerilogListener) EnterSystem_task_enable(ctx *System_task_enableContext) {}

// ExitSystem_task_enable is called when production system_task_enable is exited.
func (s *BaseVerilogListener) ExitSystem_task_enable(ctx *System_task_enableContext) {}

// EnterTask_enable is called when production task_enable is entered.
func (s *BaseVerilogListener) EnterTask_enable(ctx *Task_enableContext) {}

// ExitTask_enable is called when production task_enable is exited.
func (s *BaseVerilogListener) ExitTask_enable(ctx *Task_enableContext) {}

// EnterSpecify_block is called when production specify_block is entered.
func (s *BaseVerilogListener) EnterSpecify_block(ctx *Specify_blockContext) {}

// ExitSpecify_block is called when production specify_block is exited.
func (s *BaseVerilogListener) ExitSpecify_block(ctx *Specify_blockContext) {}

// EnterSpecify_item is called when production specify_item is entered.
func (s *BaseVerilogListener) EnterSpecify_item(ctx *Specify_itemContext) {}

// ExitSpecify_item is called when production specify_item is exited.
func (s *BaseVerilogListener) ExitSpecify_item(ctx *Specify_itemContext) {}

// EnterPulsestyle_declaration is called when production pulsestyle_declaration is entered.
func (s *BaseVerilogListener) EnterPulsestyle_declaration(ctx *Pulsestyle_declarationContext) {}

// ExitPulsestyle_declaration is called when production pulsestyle_declaration is exited.
func (s *BaseVerilogListener) ExitPulsestyle_declaration(ctx *Pulsestyle_declarationContext) {}

// EnterShowcancelled_declaration is called when production showcancelled_declaration is entered.
func (s *BaseVerilogListener) EnterShowcancelled_declaration(ctx *Showcancelled_declarationContext) {}

// ExitShowcancelled_declaration is called when production showcancelled_declaration is exited.
func (s *BaseVerilogListener) ExitShowcancelled_declaration(ctx *Showcancelled_declarationContext) {}

// EnterPath_declaration is called when production path_declaration is entered.
func (s *BaseVerilogListener) EnterPath_declaration(ctx *Path_declarationContext) {}

// ExitPath_declaration is called when production path_declaration is exited.
func (s *BaseVerilogListener) ExitPath_declaration(ctx *Path_declarationContext) {}

// EnterSimple_path_declaration is called when production simple_path_declaration is entered.
func (s *BaseVerilogListener) EnterSimple_path_declaration(ctx *Simple_path_declarationContext) {}

// ExitSimple_path_declaration is called when production simple_path_declaration is exited.
func (s *BaseVerilogListener) ExitSimple_path_declaration(ctx *Simple_path_declarationContext) {}

// EnterParallel_path_description is called when production parallel_path_description is entered.
func (s *BaseVerilogListener) EnterParallel_path_description(ctx *Parallel_path_descriptionContext) {}

// ExitParallel_path_description is called when production parallel_path_description is exited.
func (s *BaseVerilogListener) ExitParallel_path_description(ctx *Parallel_path_descriptionContext) {}

// EnterFull_path_description is called when production full_path_description is entered.
func (s *BaseVerilogListener) EnterFull_path_description(ctx *Full_path_descriptionContext) {}

// ExitFull_path_description is called when production full_path_description is exited.
func (s *BaseVerilogListener) ExitFull_path_description(ctx *Full_path_descriptionContext) {}

// EnterList_of_path_inputs is called when production list_of_path_inputs is entered.
func (s *BaseVerilogListener) EnterList_of_path_inputs(ctx *List_of_path_inputsContext) {}

// ExitList_of_path_inputs is called when production list_of_path_inputs is exited.
func (s *BaseVerilogListener) ExitList_of_path_inputs(ctx *List_of_path_inputsContext) {}

// EnterList_of_path_outputs is called when production list_of_path_outputs is entered.
func (s *BaseVerilogListener) EnterList_of_path_outputs(ctx *List_of_path_outputsContext) {}

// ExitList_of_path_outputs is called when production list_of_path_outputs is exited.
func (s *BaseVerilogListener) ExitList_of_path_outputs(ctx *List_of_path_outputsContext) {}

// EnterSpecify_input_terminal_descriptor is called when production specify_input_terminal_descriptor is entered.
func (s *BaseVerilogListener) EnterSpecify_input_terminal_descriptor(ctx *Specify_input_terminal_descriptorContext) {
}

// ExitSpecify_input_terminal_descriptor is called when production specify_input_terminal_descriptor is exited.
func (s *BaseVerilogListener) ExitSpecify_input_terminal_descriptor(ctx *Specify_input_terminal_descriptorContext) {
}

// EnterSpecify_output_terminal_descriptor is called when production specify_output_terminal_descriptor is entered.
func (s *BaseVerilogListener) EnterSpecify_output_terminal_descriptor(ctx *Specify_output_terminal_descriptorContext) {
}

// ExitSpecify_output_terminal_descriptor is called when production specify_output_terminal_descriptor is exited.
func (s *BaseVerilogListener) ExitSpecify_output_terminal_descriptor(ctx *Specify_output_terminal_descriptorContext) {
}

// EnterInput_identifier is called when production input_identifier is entered.
func (s *BaseVerilogListener) EnterInput_identifier(ctx *Input_identifierContext) {}

// ExitInput_identifier is called when production input_identifier is exited.
func (s *BaseVerilogListener) ExitInput_identifier(ctx *Input_identifierContext) {}

// EnterOutput_identifier is called when production output_identifier is entered.
func (s *BaseVerilogListener) EnterOutput_identifier(ctx *Output_identifierContext) {}

// ExitOutput_identifier is called when production output_identifier is exited.
func (s *BaseVerilogListener) ExitOutput_identifier(ctx *Output_identifierContext) {}

// EnterPath_delay_value is called when production path_delay_value is entered.
func (s *BaseVerilogListener) EnterPath_delay_value(ctx *Path_delay_valueContext) {}

// ExitPath_delay_value is called when production path_delay_value is exited.
func (s *BaseVerilogListener) ExitPath_delay_value(ctx *Path_delay_valueContext) {}

// EnterList_of_path_delay_expressions is called when production list_of_path_delay_expressions is entered.
func (s *BaseVerilogListener) EnterList_of_path_delay_expressions(ctx *List_of_path_delay_expressionsContext) {
}

// ExitList_of_path_delay_expressions is called when production list_of_path_delay_expressions is exited.
func (s *BaseVerilogListener) ExitList_of_path_delay_expressions(ctx *List_of_path_delay_expressionsContext) {
}

// EnterT_path_delay_expression is called when production t_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterT_path_delay_expression(ctx *T_path_delay_expressionContext) {}

// ExitT_path_delay_expression is called when production t_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitT_path_delay_expression(ctx *T_path_delay_expressionContext) {}

// EnterTrise_path_delay_expression is called when production trise_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterTrise_path_delay_expression(ctx *Trise_path_delay_expressionContext) {
}

// ExitTrise_path_delay_expression is called when production trise_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitTrise_path_delay_expression(ctx *Trise_path_delay_expressionContext) {
}

// EnterTfall_path_delay_expression is called when production tfall_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterTfall_path_delay_expression(ctx *Tfall_path_delay_expressionContext) {
}

// ExitTfall_path_delay_expression is called when production tfall_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitTfall_path_delay_expression(ctx *Tfall_path_delay_expressionContext) {
}

// EnterTz_path_delay_expression is called when production tz_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterTz_path_delay_expression(ctx *Tz_path_delay_expressionContext) {}

// ExitTz_path_delay_expression is called when production tz_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitTz_path_delay_expression(ctx *Tz_path_delay_expressionContext) {}

// EnterT01_path_delay_expression is called when production t01_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterT01_path_delay_expression(ctx *T01_path_delay_expressionContext) {}

// ExitT01_path_delay_expression is called when production t01_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitT01_path_delay_expression(ctx *T01_path_delay_expressionContext) {}

// EnterT10_path_delay_expression is called when production t10_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterT10_path_delay_expression(ctx *T10_path_delay_expressionContext) {}

// ExitT10_path_delay_expression is called when production t10_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitT10_path_delay_expression(ctx *T10_path_delay_expressionContext) {}

// EnterT0z_path_delay_expression is called when production t0z_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterT0z_path_delay_expression(ctx *T0z_path_delay_expressionContext) {}

// ExitT0z_path_delay_expression is called when production t0z_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitT0z_path_delay_expression(ctx *T0z_path_delay_expressionContext) {}

// EnterTz1_path_delay_expression is called when production tz1_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterTz1_path_delay_expression(ctx *Tz1_path_delay_expressionContext) {}

// ExitTz1_path_delay_expression is called when production tz1_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitTz1_path_delay_expression(ctx *Tz1_path_delay_expressionContext) {}

// EnterT1z_path_delay_expression is called when production t1z_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterT1z_path_delay_expression(ctx *T1z_path_delay_expressionContext) {}

// ExitT1z_path_delay_expression is called when production t1z_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitT1z_path_delay_expression(ctx *T1z_path_delay_expressionContext) {}

// EnterTz0_path_delay_expression is called when production tz0_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterTz0_path_delay_expression(ctx *Tz0_path_delay_expressionContext) {}

// ExitTz0_path_delay_expression is called when production tz0_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitTz0_path_delay_expression(ctx *Tz0_path_delay_expressionContext) {}

// EnterT0x_path_delay_expression is called when production t0x_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterT0x_path_delay_expression(ctx *T0x_path_delay_expressionContext) {}

// ExitT0x_path_delay_expression is called when production t0x_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitT0x_path_delay_expression(ctx *T0x_path_delay_expressionContext) {}

// EnterTx1_path_delay_expression is called when production tx1_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterTx1_path_delay_expression(ctx *Tx1_path_delay_expressionContext) {}

// ExitTx1_path_delay_expression is called when production tx1_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitTx1_path_delay_expression(ctx *Tx1_path_delay_expressionContext) {}

// EnterT1x_path_delay_expression is called when production t1x_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterT1x_path_delay_expression(ctx *T1x_path_delay_expressionContext) {}

// ExitT1x_path_delay_expression is called when production t1x_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitT1x_path_delay_expression(ctx *T1x_path_delay_expressionContext) {}

// EnterTx0_path_delay_expression is called when production tx0_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterTx0_path_delay_expression(ctx *Tx0_path_delay_expressionContext) {}

// ExitTx0_path_delay_expression is called when production tx0_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitTx0_path_delay_expression(ctx *Tx0_path_delay_expressionContext) {}

// EnterTxz_path_delay_expression is called when production txz_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterTxz_path_delay_expression(ctx *Txz_path_delay_expressionContext) {}

// ExitTxz_path_delay_expression is called when production txz_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitTxz_path_delay_expression(ctx *Txz_path_delay_expressionContext) {}

// EnterTzx_path_delay_expression is called when production tzx_path_delay_expression is entered.
func (s *BaseVerilogListener) EnterTzx_path_delay_expression(ctx *Tzx_path_delay_expressionContext) {}

// ExitTzx_path_delay_expression is called when production tzx_path_delay_expression is exited.
func (s *BaseVerilogListener) ExitTzx_path_delay_expression(ctx *Tzx_path_delay_expressionContext) {}

// EnterPath_delay_expression is called when production path_delay_expression is entered.
func (s *BaseVerilogListener) EnterPath_delay_expression(ctx *Path_delay_expressionContext) {}

// ExitPath_delay_expression is called when production path_delay_expression is exited.
func (s *BaseVerilogListener) ExitPath_delay_expression(ctx *Path_delay_expressionContext) {}

// EnterEdge_sensitive_path_declaration is called when production edge_sensitive_path_declaration is entered.
func (s *BaseVerilogListener) EnterEdge_sensitive_path_declaration(ctx *Edge_sensitive_path_declarationContext) {
}

// ExitEdge_sensitive_path_declaration is called when production edge_sensitive_path_declaration is exited.
func (s *BaseVerilogListener) ExitEdge_sensitive_path_declaration(ctx *Edge_sensitive_path_declarationContext) {
}

// EnterParallel_edge_sensitive_path_description is called when production parallel_edge_sensitive_path_description is entered.
func (s *BaseVerilogListener) EnterParallel_edge_sensitive_path_description(ctx *Parallel_edge_sensitive_path_descriptionContext) {
}

// ExitParallel_edge_sensitive_path_description is called when production parallel_edge_sensitive_path_description is exited.
func (s *BaseVerilogListener) ExitParallel_edge_sensitive_path_description(ctx *Parallel_edge_sensitive_path_descriptionContext) {
}

// EnterFull_edge_sensitive_path_description is called when production full_edge_sensitive_path_description is entered.
func (s *BaseVerilogListener) EnterFull_edge_sensitive_path_description(ctx *Full_edge_sensitive_path_descriptionContext) {
}

// ExitFull_edge_sensitive_path_description is called when production full_edge_sensitive_path_description is exited.
func (s *BaseVerilogListener) ExitFull_edge_sensitive_path_description(ctx *Full_edge_sensitive_path_descriptionContext) {
}

// EnterData_source_expression is called when production data_source_expression is entered.
func (s *BaseVerilogListener) EnterData_source_expression(ctx *Data_source_expressionContext) {}

// ExitData_source_expression is called when production data_source_expression is exited.
func (s *BaseVerilogListener) ExitData_source_expression(ctx *Data_source_expressionContext) {}

// EnterEdge_identifier is called when production edge_identifier is entered.
func (s *BaseVerilogListener) EnterEdge_identifier(ctx *Edge_identifierContext) {}

// ExitEdge_identifier is called when production edge_identifier is exited.
func (s *BaseVerilogListener) ExitEdge_identifier(ctx *Edge_identifierContext) {}

// EnterState_dependent_path_declaration is called when production state_dependent_path_declaration is entered.
func (s *BaseVerilogListener) EnterState_dependent_path_declaration(ctx *State_dependent_path_declarationContext) {
}

// ExitState_dependent_path_declaration is called when production state_dependent_path_declaration is exited.
func (s *BaseVerilogListener) ExitState_dependent_path_declaration(ctx *State_dependent_path_declarationContext) {
}

// EnterPolarity_operator is called when production polarity_operator is entered.
func (s *BaseVerilogListener) EnterPolarity_operator(ctx *Polarity_operatorContext) {}

// ExitPolarity_operator is called when production polarity_operator is exited.
func (s *BaseVerilogListener) ExitPolarity_operator(ctx *Polarity_operatorContext) {}

// EnterChecktime_condition is called when production checktime_condition is entered.
func (s *BaseVerilogListener) EnterChecktime_condition(ctx *Checktime_conditionContext) {}

// ExitChecktime_condition is called when production checktime_condition is exited.
func (s *BaseVerilogListener) ExitChecktime_condition(ctx *Checktime_conditionContext) {}

// EnterDelayed_data is called when production delayed_data is entered.
func (s *BaseVerilogListener) EnterDelayed_data(ctx *Delayed_dataContext) {}

// ExitDelayed_data is called when production delayed_data is exited.
func (s *BaseVerilogListener) ExitDelayed_data(ctx *Delayed_dataContext) {}

// EnterDelayed_reference is called when production delayed_reference is entered.
func (s *BaseVerilogListener) EnterDelayed_reference(ctx *Delayed_referenceContext) {}

// ExitDelayed_reference is called when production delayed_reference is exited.
func (s *BaseVerilogListener) ExitDelayed_reference(ctx *Delayed_referenceContext) {}

// EnterEnd_edge_offset is called when production end_edge_offset is entered.
func (s *BaseVerilogListener) EnterEnd_edge_offset(ctx *End_edge_offsetContext) {}

// ExitEnd_edge_offset is called when production end_edge_offset is exited.
func (s *BaseVerilogListener) ExitEnd_edge_offset(ctx *End_edge_offsetContext) {}

// EnterEvent_based_flag is called when production event_based_flag is entered.
func (s *BaseVerilogListener) EnterEvent_based_flag(ctx *Event_based_flagContext) {}

// ExitEvent_based_flag is called when production event_based_flag is exited.
func (s *BaseVerilogListener) ExitEvent_based_flag(ctx *Event_based_flagContext) {}

// EnterNotify_reg is called when production notify_reg is entered.
func (s *BaseVerilogListener) EnterNotify_reg(ctx *Notify_regContext) {}

// ExitNotify_reg is called when production notify_reg is exited.
func (s *BaseVerilogListener) ExitNotify_reg(ctx *Notify_regContext) {}

// EnterRemain_active_flag is called when production remain_active_flag is entered.
func (s *BaseVerilogListener) EnterRemain_active_flag(ctx *Remain_active_flagContext) {}

// ExitRemain_active_flag is called when production remain_active_flag is exited.
func (s *BaseVerilogListener) ExitRemain_active_flag(ctx *Remain_active_flagContext) {}

// EnterStamptime_condition is called when production stamptime_condition is entered.
func (s *BaseVerilogListener) EnterStamptime_condition(ctx *Stamptime_conditionContext) {}

// ExitStamptime_condition is called when production stamptime_condition is exited.
func (s *BaseVerilogListener) ExitStamptime_condition(ctx *Stamptime_conditionContext) {}

// EnterStart_edge_offset is called when production start_edge_offset is entered.
func (s *BaseVerilogListener) EnterStart_edge_offset(ctx *Start_edge_offsetContext) {}

// ExitStart_edge_offset is called when production start_edge_offset is exited.
func (s *BaseVerilogListener) ExitStart_edge_offset(ctx *Start_edge_offsetContext) {}

// EnterThreshold is called when production threshold is entered.
func (s *BaseVerilogListener) EnterThreshold(ctx *ThresholdContext) {}

// ExitThreshold is called when production threshold is exited.
func (s *BaseVerilogListener) ExitThreshold(ctx *ThresholdContext) {}

// EnterTiming_check_limit is called when production timing_check_limit is entered.
func (s *BaseVerilogListener) EnterTiming_check_limit(ctx *Timing_check_limitContext) {}

// ExitTiming_check_limit is called when production timing_check_limit is exited.
func (s *BaseVerilogListener) ExitTiming_check_limit(ctx *Timing_check_limitContext) {}

// EnterConcatenation is called when production concatenation is entered.
func (s *BaseVerilogListener) EnterConcatenation(ctx *ConcatenationContext) {}

// ExitConcatenation is called when production concatenation is exited.
func (s *BaseVerilogListener) ExitConcatenation(ctx *ConcatenationContext) {}

// EnterConstant_concatenation is called when production constant_concatenation is entered.
func (s *BaseVerilogListener) EnterConstant_concatenation(ctx *Constant_concatenationContext) {}

// ExitConstant_concatenation is called when production constant_concatenation is exited.
func (s *BaseVerilogListener) ExitConstant_concatenation(ctx *Constant_concatenationContext) {}

// EnterConstant_multiple_concatenation is called when production constant_multiple_concatenation is entered.
func (s *BaseVerilogListener) EnterConstant_multiple_concatenation(ctx *Constant_multiple_concatenationContext) {
}

// ExitConstant_multiple_concatenation is called when production constant_multiple_concatenation is exited.
func (s *BaseVerilogListener) ExitConstant_multiple_concatenation(ctx *Constant_multiple_concatenationContext) {
}

// EnterModule_path_concatenation is called when production module_path_concatenation is entered.
func (s *BaseVerilogListener) EnterModule_path_concatenation(ctx *Module_path_concatenationContext) {}

// ExitModule_path_concatenation is called when production module_path_concatenation is exited.
func (s *BaseVerilogListener) ExitModule_path_concatenation(ctx *Module_path_concatenationContext) {}

// EnterModule_path_multiple_concatenation is called when production module_path_multiple_concatenation is entered.
func (s *BaseVerilogListener) EnterModule_path_multiple_concatenation(ctx *Module_path_multiple_concatenationContext) {
}

// ExitModule_path_multiple_concatenation is called when production module_path_multiple_concatenation is exited.
func (s *BaseVerilogListener) ExitModule_path_multiple_concatenation(ctx *Module_path_multiple_concatenationContext) {
}

// EnterMultiple_concatenation is called when production multiple_concatenation is entered.
func (s *BaseVerilogListener) EnterMultiple_concatenation(ctx *Multiple_concatenationContext) {}

// ExitMultiple_concatenation is called when production multiple_concatenation is exited.
func (s *BaseVerilogListener) ExitMultiple_concatenation(ctx *Multiple_concatenationContext) {}

// EnterNet_concatenation is called when production net_concatenation is entered.
func (s *BaseVerilogListener) EnterNet_concatenation(ctx *Net_concatenationContext) {}

// ExitNet_concatenation is called when production net_concatenation is exited.
func (s *BaseVerilogListener) ExitNet_concatenation(ctx *Net_concatenationContext) {}

// EnterNet_concatenation_value is called when production net_concatenation_value is entered.
func (s *BaseVerilogListener) EnterNet_concatenation_value(ctx *Net_concatenation_valueContext) {}

// ExitNet_concatenation_value is called when production net_concatenation_value is exited.
func (s *BaseVerilogListener) ExitNet_concatenation_value(ctx *Net_concatenation_valueContext) {}

// EnterVariable_concatenation is called when production variable_concatenation is entered.
func (s *BaseVerilogListener) EnterVariable_concatenation(ctx *Variable_concatenationContext) {}

// ExitVariable_concatenation is called when production variable_concatenation is exited.
func (s *BaseVerilogListener) ExitVariable_concatenation(ctx *Variable_concatenationContext) {}

// EnterVariable_concatenation_value is called when production variable_concatenation_value is entered.
func (s *BaseVerilogListener) EnterVariable_concatenation_value(ctx *Variable_concatenation_valueContext) {
}

// ExitVariable_concatenation_value is called when production variable_concatenation_value is exited.
func (s *BaseVerilogListener) ExitVariable_concatenation_value(ctx *Variable_concatenation_valueContext) {
}

// EnterConstant_function_call is called when production constant_function_call is entered.
func (s *BaseVerilogListener) EnterConstant_function_call(ctx *Constant_function_callContext) {}

// ExitConstant_function_call is called when production constant_function_call is exited.
func (s *BaseVerilogListener) ExitConstant_function_call(ctx *Constant_function_callContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseVerilogListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseVerilogListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterSystem_function_call is called when production system_function_call is entered.
func (s *BaseVerilogListener) EnterSystem_function_call(ctx *System_function_callContext) {}

// ExitSystem_function_call is called when production system_function_call is exited.
func (s *BaseVerilogListener) ExitSystem_function_call(ctx *System_function_callContext) {}

// EnterGenvar_function_call is called when production genvar_function_call is entered.
func (s *BaseVerilogListener) EnterGenvar_function_call(ctx *Genvar_function_callContext) {}

// ExitGenvar_function_call is called when production genvar_function_call is exited.
func (s *BaseVerilogListener) ExitGenvar_function_call(ctx *Genvar_function_callContext) {}

// EnterBase_expression is called when production base_expression is entered.
func (s *BaseVerilogListener) EnterBase_expression(ctx *Base_expressionContext) {}

// ExitBase_expression is called when production base_expression is exited.
func (s *BaseVerilogListener) ExitBase_expression(ctx *Base_expressionContext) {}

// EnterConstant_base_expression is called when production constant_base_expression is entered.
func (s *BaseVerilogListener) EnterConstant_base_expression(ctx *Constant_base_expressionContext) {}

// ExitConstant_base_expression is called when production constant_base_expression is exited.
func (s *BaseVerilogListener) ExitConstant_base_expression(ctx *Constant_base_expressionContext) {}

// EnterConstant_expression is called when production constant_expression is entered.
func (s *BaseVerilogListener) EnterConstant_expression(ctx *Constant_expressionContext) {}

// ExitConstant_expression is called when production constant_expression is exited.
func (s *BaseVerilogListener) ExitConstant_expression(ctx *Constant_expressionContext) {}

// EnterConstant_mintypmax_expression is called when production constant_mintypmax_expression is entered.
func (s *BaseVerilogListener) EnterConstant_mintypmax_expression(ctx *Constant_mintypmax_expressionContext) {
}

// ExitConstant_mintypmax_expression is called when production constant_mintypmax_expression is exited.
func (s *BaseVerilogListener) ExitConstant_mintypmax_expression(ctx *Constant_mintypmax_expressionContext) {
}

// EnterConstant_range_expression is called when production constant_range_expression is entered.
func (s *BaseVerilogListener) EnterConstant_range_expression(ctx *Constant_range_expressionContext) {}

// ExitConstant_range_expression is called when production constant_range_expression is exited.
func (s *BaseVerilogListener) ExitConstant_range_expression(ctx *Constant_range_expressionContext) {}

// EnterDimension_constant_expression is called when production dimension_constant_expression is entered.
func (s *BaseVerilogListener) EnterDimension_constant_expression(ctx *Dimension_constant_expressionContext) {
}

// ExitDimension_constant_expression is called when production dimension_constant_expression is exited.
func (s *BaseVerilogListener) ExitDimension_constant_expression(ctx *Dimension_constant_expressionContext) {
}

// EnterExpression is called when production expression is entered.
func (s *BaseVerilogListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVerilogListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseVerilogListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseVerilogListener) ExitTerm(ctx *TermContext) {}

// EnterLsb_constant_expression is called when production lsb_constant_expression is entered.
func (s *BaseVerilogListener) EnterLsb_constant_expression(ctx *Lsb_constant_expressionContext) {}

// ExitLsb_constant_expression is called when production lsb_constant_expression is exited.
func (s *BaseVerilogListener) ExitLsb_constant_expression(ctx *Lsb_constant_expressionContext) {}

// EnterMintypmax_expression is called when production mintypmax_expression is entered.
func (s *BaseVerilogListener) EnterMintypmax_expression(ctx *Mintypmax_expressionContext) {}

// ExitMintypmax_expression is called when production mintypmax_expression is exited.
func (s *BaseVerilogListener) ExitMintypmax_expression(ctx *Mintypmax_expressionContext) {}

// EnterModule_path_conditional_expression is called when production module_path_conditional_expression is entered.
func (s *BaseVerilogListener) EnterModule_path_conditional_expression(ctx *Module_path_conditional_expressionContext) {
}

// ExitModule_path_conditional_expression is called when production module_path_conditional_expression is exited.
func (s *BaseVerilogListener) ExitModule_path_conditional_expression(ctx *Module_path_conditional_expressionContext) {
}

// EnterModule_path_expression is called when production module_path_expression is entered.
func (s *BaseVerilogListener) EnterModule_path_expression(ctx *Module_path_expressionContext) {}

// ExitModule_path_expression is called when production module_path_expression is exited.
func (s *BaseVerilogListener) ExitModule_path_expression(ctx *Module_path_expressionContext) {}

// EnterModule_path_mintypmax_expression is called when production module_path_mintypmax_expression is entered.
func (s *BaseVerilogListener) EnterModule_path_mintypmax_expression(ctx *Module_path_mintypmax_expressionContext) {
}

// ExitModule_path_mintypmax_expression is called when production module_path_mintypmax_expression is exited.
func (s *BaseVerilogListener) ExitModule_path_mintypmax_expression(ctx *Module_path_mintypmax_expressionContext) {
}

// EnterMsb_constant_expression is called when production msb_constant_expression is entered.
func (s *BaseVerilogListener) EnterMsb_constant_expression(ctx *Msb_constant_expressionContext) {}

// ExitMsb_constant_expression is called when production msb_constant_expression is exited.
func (s *BaseVerilogListener) ExitMsb_constant_expression(ctx *Msb_constant_expressionContext) {}

// EnterRange_expression is called when production range_expression is entered.
func (s *BaseVerilogListener) EnterRange_expression(ctx *Range_expressionContext) {}

// ExitRange_expression is called when production range_expression is exited.
func (s *BaseVerilogListener) ExitRange_expression(ctx *Range_expressionContext) {}

// EnterWidth_constant_expression is called when production width_constant_expression is entered.
func (s *BaseVerilogListener) EnterWidth_constant_expression(ctx *Width_constant_expressionContext) {}

// ExitWidth_constant_expression is called when production width_constant_expression is exited.
func (s *BaseVerilogListener) ExitWidth_constant_expression(ctx *Width_constant_expressionContext) {}

// EnterConstant_primary is called when production constant_primary is entered.
func (s *BaseVerilogListener) EnterConstant_primary(ctx *Constant_primaryContext) {}

// ExitConstant_primary is called when production constant_primary is exited.
func (s *BaseVerilogListener) ExitConstant_primary(ctx *Constant_primaryContext) {}

// EnterModule_path_primary is called when production module_path_primary is entered.
func (s *BaseVerilogListener) EnterModule_path_primary(ctx *Module_path_primaryContext) {}

// ExitModule_path_primary is called when production module_path_primary is exited.
func (s *BaseVerilogListener) ExitModule_path_primary(ctx *Module_path_primaryContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BaseVerilogListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BaseVerilogListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterNet_lvalue is called when production net_lvalue is entered.
func (s *BaseVerilogListener) EnterNet_lvalue(ctx *Net_lvalueContext) {}

// ExitNet_lvalue is called when production net_lvalue is exited.
func (s *BaseVerilogListener) ExitNet_lvalue(ctx *Net_lvalueContext) {}

// EnterVariable_lvalue is called when production variable_lvalue is entered.
func (s *BaseVerilogListener) EnterVariable_lvalue(ctx *Variable_lvalueContext) {}

// ExitVariable_lvalue is called when production variable_lvalue is exited.
func (s *BaseVerilogListener) ExitVariable_lvalue(ctx *Variable_lvalueContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseVerilogListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseVerilogListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterBinary_operator is called when production binary_operator is entered.
func (s *BaseVerilogListener) EnterBinary_operator(ctx *Binary_operatorContext) {}

// ExitBinary_operator is called when production binary_operator is exited.
func (s *BaseVerilogListener) ExitBinary_operator(ctx *Binary_operatorContext) {}

// EnterUnary_module_path_operator is called when production unary_module_path_operator is entered.
func (s *BaseVerilogListener) EnterUnary_module_path_operator(ctx *Unary_module_path_operatorContext) {
}

// ExitUnary_module_path_operator is called when production unary_module_path_operator is exited.
func (s *BaseVerilogListener) ExitUnary_module_path_operator(ctx *Unary_module_path_operatorContext) {
}

// EnterBinary_module_path_operator is called when production binary_module_path_operator is entered.
func (s *BaseVerilogListener) EnterBinary_module_path_operator(ctx *Binary_module_path_operatorContext) {
}

// ExitBinary_module_path_operator is called when production binary_module_path_operator is exited.
func (s *BaseVerilogListener) ExitBinary_module_path_operator(ctx *Binary_module_path_operatorContext) {
}

// EnterNumber is called when production number is entered.
func (s *BaseVerilogListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseVerilogListener) ExitNumber(ctx *NumberContext) {}

// EnterTiming_spec is called when production timing_spec is entered.
func (s *BaseVerilogListener) EnterTiming_spec(ctx *Timing_specContext) {}

// ExitTiming_spec is called when production timing_spec is exited.
func (s *BaseVerilogListener) ExitTiming_spec(ctx *Timing_specContext) {}

// EnterDefine is called when production define is entered.
func (s *BaseVerilogListener) EnterDefine(ctx *DefineContext) {}

// ExitDefine is called when production define is exited.
func (s *BaseVerilogListener) ExitDefine(ctx *DefineContext) {}

// EnterAttribute_instance is called when production attribute_instance is entered.
func (s *BaseVerilogListener) EnterAttribute_instance(ctx *Attribute_instanceContext) {}

// ExitAttribute_instance is called when production attribute_instance is exited.
func (s *BaseVerilogListener) ExitAttribute_instance(ctx *Attribute_instanceContext) {}

// EnterAttr_spec is called when production attr_spec is entered.
func (s *BaseVerilogListener) EnterAttr_spec(ctx *Attr_specContext) {}

// ExitAttr_spec is called when production attr_spec is exited.
func (s *BaseVerilogListener) ExitAttr_spec(ctx *Attr_specContext) {}

// EnterAttr_name is called when production attr_name is entered.
func (s *BaseVerilogListener) EnterAttr_name(ctx *Attr_nameContext) {}

// ExitAttr_name is called when production attr_name is exited.
func (s *BaseVerilogListener) ExitAttr_name(ctx *Attr_nameContext) {}

// EnterArrayed_identifier is called when production arrayed_identifier is entered.
func (s *BaseVerilogListener) EnterArrayed_identifier(ctx *Arrayed_identifierContext) {}

// ExitArrayed_identifier is called when production arrayed_identifier is exited.
func (s *BaseVerilogListener) ExitArrayed_identifier(ctx *Arrayed_identifierContext) {}

// EnterBlock_identifier is called when production block_identifier is entered.
func (s *BaseVerilogListener) EnterBlock_identifier(ctx *Block_identifierContext) {}

// ExitBlock_identifier is called when production block_identifier is exited.
func (s *BaseVerilogListener) ExitBlock_identifier(ctx *Block_identifierContext) {}

// EnterCell_identifier is called when production cell_identifier is entered.
func (s *BaseVerilogListener) EnterCell_identifier(ctx *Cell_identifierContext) {}

// ExitCell_identifier is called when production cell_identifier is exited.
func (s *BaseVerilogListener) ExitCell_identifier(ctx *Cell_identifierContext) {}

// EnterConfig_identifier is called when production config_identifier is entered.
func (s *BaseVerilogListener) EnterConfig_identifier(ctx *Config_identifierContext) {}

// ExitConfig_identifier is called when production config_identifier is exited.
func (s *BaseVerilogListener) ExitConfig_identifier(ctx *Config_identifierContext) {}

// EnterDefine_identifier is called when production define_identifier is entered.
func (s *BaseVerilogListener) EnterDefine_identifier(ctx *Define_identifierContext) {}

// ExitDefine_identifier is called when production define_identifier is exited.
func (s *BaseVerilogListener) ExitDefine_identifier(ctx *Define_identifierContext) {}

// EnterEscaped_arrayed_identifier is called when production escaped_arrayed_identifier is entered.
func (s *BaseVerilogListener) EnterEscaped_arrayed_identifier(ctx *Escaped_arrayed_identifierContext) {
}

// ExitEscaped_arrayed_identifier is called when production escaped_arrayed_identifier is exited.
func (s *BaseVerilogListener) ExitEscaped_arrayed_identifier(ctx *Escaped_arrayed_identifierContext) {
}

// EnterEscaped_hierarchical_identifier is called when production escaped_hierarchical_identifier is entered.
func (s *BaseVerilogListener) EnterEscaped_hierarchical_identifier(ctx *Escaped_hierarchical_identifierContext) {
}

// ExitEscaped_hierarchical_identifier is called when production escaped_hierarchical_identifier is exited.
func (s *BaseVerilogListener) ExitEscaped_hierarchical_identifier(ctx *Escaped_hierarchical_identifierContext) {
}

// EnterEvent_identifier is called when production event_identifier is entered.
func (s *BaseVerilogListener) EnterEvent_identifier(ctx *Event_identifierContext) {}

// ExitEvent_identifier is called when production event_identifier is exited.
func (s *BaseVerilogListener) ExitEvent_identifier(ctx *Event_identifierContext) {}

// EnterFunction_identifier is called when production function_identifier is entered.
func (s *BaseVerilogListener) EnterFunction_identifier(ctx *Function_identifierContext) {}

// ExitFunction_identifier is called when production function_identifier is exited.
func (s *BaseVerilogListener) ExitFunction_identifier(ctx *Function_identifierContext) {}

// EnterGate_instance_identifier is called when production gate_instance_identifier is entered.
func (s *BaseVerilogListener) EnterGate_instance_identifier(ctx *Gate_instance_identifierContext) {}

// ExitGate_instance_identifier is called when production gate_instance_identifier is exited.
func (s *BaseVerilogListener) ExitGate_instance_identifier(ctx *Gate_instance_identifierContext) {}

// EnterGenerate_block_identifier is called when production generate_block_identifier is entered.
func (s *BaseVerilogListener) EnterGenerate_block_identifier(ctx *Generate_block_identifierContext) {}

// ExitGenerate_block_identifier is called when production generate_block_identifier is exited.
func (s *BaseVerilogListener) ExitGenerate_block_identifier(ctx *Generate_block_identifierContext) {}

// EnterGenvar_function_identifier is called when production genvar_function_identifier is entered.
func (s *BaseVerilogListener) EnterGenvar_function_identifier(ctx *Genvar_function_identifierContext) {
}

// ExitGenvar_function_identifier is called when production genvar_function_identifier is exited.
func (s *BaseVerilogListener) ExitGenvar_function_identifier(ctx *Genvar_function_identifierContext) {
}

// EnterGenvar_identifier is called when production genvar_identifier is entered.
func (s *BaseVerilogListener) EnterGenvar_identifier(ctx *Genvar_identifierContext) {}

// ExitGenvar_identifier is called when production genvar_identifier is exited.
func (s *BaseVerilogListener) ExitGenvar_identifier(ctx *Genvar_identifierContext) {}

// EnterHierarchical_block_identifier is called when production hierarchical_block_identifier is entered.
func (s *BaseVerilogListener) EnterHierarchical_block_identifier(ctx *Hierarchical_block_identifierContext) {
}

// ExitHierarchical_block_identifier is called when production hierarchical_block_identifier is exited.
func (s *BaseVerilogListener) ExitHierarchical_block_identifier(ctx *Hierarchical_block_identifierContext) {
}

// EnterHierarchical_event_identifier is called when production hierarchical_event_identifier is entered.
func (s *BaseVerilogListener) EnterHierarchical_event_identifier(ctx *Hierarchical_event_identifierContext) {
}

// ExitHierarchical_event_identifier is called when production hierarchical_event_identifier is exited.
func (s *BaseVerilogListener) ExitHierarchical_event_identifier(ctx *Hierarchical_event_identifierContext) {
}

// EnterHierarchical_function_identifier is called when production hierarchical_function_identifier is entered.
func (s *BaseVerilogListener) EnterHierarchical_function_identifier(ctx *Hierarchical_function_identifierContext) {
}

// ExitHierarchical_function_identifier is called when production hierarchical_function_identifier is exited.
func (s *BaseVerilogListener) ExitHierarchical_function_identifier(ctx *Hierarchical_function_identifierContext) {
}

// EnterHierarchical_identifier is called when production hierarchical_identifier is entered.
func (s *BaseVerilogListener) EnterHierarchical_identifier(ctx *Hierarchical_identifierContext) {}

// ExitHierarchical_identifier is called when production hierarchical_identifier is exited.
func (s *BaseVerilogListener) ExitHierarchical_identifier(ctx *Hierarchical_identifierContext) {}

// EnterHierarchical_net_identifier is called when production hierarchical_net_identifier is entered.
func (s *BaseVerilogListener) EnterHierarchical_net_identifier(ctx *Hierarchical_net_identifierContext) {
}

// ExitHierarchical_net_identifier is called when production hierarchical_net_identifier is exited.
func (s *BaseVerilogListener) ExitHierarchical_net_identifier(ctx *Hierarchical_net_identifierContext) {
}

// EnterHierarchical_variable_identifier is called when production hierarchical_variable_identifier is entered.
func (s *BaseVerilogListener) EnterHierarchical_variable_identifier(ctx *Hierarchical_variable_identifierContext) {
}

// ExitHierarchical_variable_identifier is called when production hierarchical_variable_identifier is exited.
func (s *BaseVerilogListener) ExitHierarchical_variable_identifier(ctx *Hierarchical_variable_identifierContext) {
}

// EnterHierarchical_task_identifier is called when production hierarchical_task_identifier is entered.
func (s *BaseVerilogListener) EnterHierarchical_task_identifier(ctx *Hierarchical_task_identifierContext) {
}

// ExitHierarchical_task_identifier is called when production hierarchical_task_identifier is exited.
func (s *BaseVerilogListener) ExitHierarchical_task_identifier(ctx *Hierarchical_task_identifierContext) {
}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseVerilogListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseVerilogListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterInout_port_identifier is called when production inout_port_identifier is entered.
func (s *BaseVerilogListener) EnterInout_port_identifier(ctx *Inout_port_identifierContext) {}

// ExitInout_port_identifier is called when production inout_port_identifier is exited.
func (s *BaseVerilogListener) ExitInout_port_identifier(ctx *Inout_port_identifierContext) {}

// EnterInput_port_identifier is called when production input_port_identifier is entered.
func (s *BaseVerilogListener) EnterInput_port_identifier(ctx *Input_port_identifierContext) {}

// ExitInput_port_identifier is called when production input_port_identifier is exited.
func (s *BaseVerilogListener) ExitInput_port_identifier(ctx *Input_port_identifierContext) {}

// EnterInstance_identifier is called when production instance_identifier is entered.
func (s *BaseVerilogListener) EnterInstance_identifier(ctx *Instance_identifierContext) {}

// ExitInstance_identifier is called when production instance_identifier is exited.
func (s *BaseVerilogListener) ExitInstance_identifier(ctx *Instance_identifierContext) {}

// EnterLibrary_identifier is called when production library_identifier is entered.
func (s *BaseVerilogListener) EnterLibrary_identifier(ctx *Library_identifierContext) {}

// ExitLibrary_identifier is called when production library_identifier is exited.
func (s *BaseVerilogListener) ExitLibrary_identifier(ctx *Library_identifierContext) {}

// EnterMemory_identifier is called when production memory_identifier is entered.
func (s *BaseVerilogListener) EnterMemory_identifier(ctx *Memory_identifierContext) {}

// ExitMemory_identifier is called when production memory_identifier is exited.
func (s *BaseVerilogListener) ExitMemory_identifier(ctx *Memory_identifierContext) {}

// EnterModule_identifier is called when production module_identifier is entered.
func (s *BaseVerilogListener) EnterModule_identifier(ctx *Module_identifierContext) {}

// ExitModule_identifier is called when production module_identifier is exited.
func (s *BaseVerilogListener) ExitModule_identifier(ctx *Module_identifierContext) {}

// EnterModule_instance_identifier is called when production module_instance_identifier is entered.
func (s *BaseVerilogListener) EnterModule_instance_identifier(ctx *Module_instance_identifierContext) {
}

// ExitModule_instance_identifier is called when production module_instance_identifier is exited.
func (s *BaseVerilogListener) ExitModule_instance_identifier(ctx *Module_instance_identifierContext) {
}

// EnterNet_identifier is called when production net_identifier is entered.
func (s *BaseVerilogListener) EnterNet_identifier(ctx *Net_identifierContext) {}

// ExitNet_identifier is called when production net_identifier is exited.
func (s *BaseVerilogListener) ExitNet_identifier(ctx *Net_identifierContext) {}

// EnterOutput_port_identifier is called when production output_port_identifier is entered.
func (s *BaseVerilogListener) EnterOutput_port_identifier(ctx *Output_port_identifierContext) {}

// ExitOutput_port_identifier is called when production output_port_identifier is exited.
func (s *BaseVerilogListener) ExitOutput_port_identifier(ctx *Output_port_identifierContext) {}

// EnterParameter_identifier is called when production parameter_identifier is entered.
func (s *BaseVerilogListener) EnterParameter_identifier(ctx *Parameter_identifierContext) {}

// ExitParameter_identifier is called when production parameter_identifier is exited.
func (s *BaseVerilogListener) ExitParameter_identifier(ctx *Parameter_identifierContext) {}

// EnterPort_identifier is called when production port_identifier is entered.
func (s *BaseVerilogListener) EnterPort_identifier(ctx *Port_identifierContext) {}

// ExitPort_identifier is called when production port_identifier is exited.
func (s *BaseVerilogListener) ExitPort_identifier(ctx *Port_identifierContext) {}

// EnterReal_identifier is called when production real_identifier is entered.
func (s *BaseVerilogListener) EnterReal_identifier(ctx *Real_identifierContext) {}

// ExitReal_identifier is called when production real_identifier is exited.
func (s *BaseVerilogListener) ExitReal_identifier(ctx *Real_identifierContext) {}

// EnterSimple_arrayed_identifier is called when production simple_arrayed_identifier is entered.
func (s *BaseVerilogListener) EnterSimple_arrayed_identifier(ctx *Simple_arrayed_identifierContext) {}

// ExitSimple_arrayed_identifier is called when production simple_arrayed_identifier is exited.
func (s *BaseVerilogListener) ExitSimple_arrayed_identifier(ctx *Simple_arrayed_identifierContext) {}

// EnterSimple_hierarchical_identifier is called when production simple_hierarchical_identifier is entered.
func (s *BaseVerilogListener) EnterSimple_hierarchical_identifier(ctx *Simple_hierarchical_identifierContext) {
}

// ExitSimple_hierarchical_identifier is called when production simple_hierarchical_identifier is exited.
func (s *BaseVerilogListener) ExitSimple_hierarchical_identifier(ctx *Simple_hierarchical_identifierContext) {
}

// EnterSpecparam_identifier is called when production specparam_identifier is entered.
func (s *BaseVerilogListener) EnterSpecparam_identifier(ctx *Specparam_identifierContext) {}

// ExitSpecparam_identifier is called when production specparam_identifier is exited.
func (s *BaseVerilogListener) ExitSpecparam_identifier(ctx *Specparam_identifierContext) {}

// EnterSystem_function_identifier is called when production system_function_identifier is entered.
func (s *BaseVerilogListener) EnterSystem_function_identifier(ctx *System_function_identifierContext) {
}

// ExitSystem_function_identifier is called when production system_function_identifier is exited.
func (s *BaseVerilogListener) ExitSystem_function_identifier(ctx *System_function_identifierContext) {
}

// EnterSystem_task_identifier is called when production system_task_identifier is entered.
func (s *BaseVerilogListener) EnterSystem_task_identifier(ctx *System_task_identifierContext) {}

// ExitSystem_task_identifier is called when production system_task_identifier is exited.
func (s *BaseVerilogListener) ExitSystem_task_identifier(ctx *System_task_identifierContext) {}

// EnterTask_identifier is called when production task_identifier is entered.
func (s *BaseVerilogListener) EnterTask_identifier(ctx *Task_identifierContext) {}

// ExitTask_identifier is called when production task_identifier is exited.
func (s *BaseVerilogListener) ExitTask_identifier(ctx *Task_identifierContext) {}

// EnterTerminal_identifier is called when production terminal_identifier is entered.
func (s *BaseVerilogListener) EnterTerminal_identifier(ctx *Terminal_identifierContext) {}

// ExitTerminal_identifier is called when production terminal_identifier is exited.
func (s *BaseVerilogListener) ExitTerminal_identifier(ctx *Terminal_identifierContext) {}

// EnterText_macro_identifier is called when production text_macro_identifier is entered.
func (s *BaseVerilogListener) EnterText_macro_identifier(ctx *Text_macro_identifierContext) {}

// ExitText_macro_identifier is called when production text_macro_identifier is exited.
func (s *BaseVerilogListener) ExitText_macro_identifier(ctx *Text_macro_identifierContext) {}

// EnterTopmodule_identifier is called when production topmodule_identifier is entered.
func (s *BaseVerilogListener) EnterTopmodule_identifier(ctx *Topmodule_identifierContext) {}

// ExitTopmodule_identifier is called when production topmodule_identifier is exited.
func (s *BaseVerilogListener) ExitTopmodule_identifier(ctx *Topmodule_identifierContext) {}

// EnterUdp_identifier is called when production udp_identifier is entered.
func (s *BaseVerilogListener) EnterUdp_identifier(ctx *Udp_identifierContext) {}

// ExitUdp_identifier is called when production udp_identifier is exited.
func (s *BaseVerilogListener) ExitUdp_identifier(ctx *Udp_identifierContext) {}

// EnterUdp_instance_identifier is called when production udp_instance_identifier is entered.
func (s *BaseVerilogListener) EnterUdp_instance_identifier(ctx *Udp_instance_identifierContext) {}

// ExitUdp_instance_identifier is called when production udp_instance_identifier is exited.
func (s *BaseVerilogListener) ExitUdp_instance_identifier(ctx *Udp_instance_identifierContext) {}

// EnterVariable_identifier is called when production variable_identifier is entered.
func (s *BaseVerilogListener) EnterVariable_identifier(ctx *Variable_identifierContext) {}

// ExitVariable_identifier is called when production variable_identifier is exited.
func (s *BaseVerilogListener) ExitVariable_identifier(ctx *Variable_identifierContext) {}

// EnterSimple_hierarchical_branch is called when production simple_hierarchical_branch is entered.
func (s *BaseVerilogListener) EnterSimple_hierarchical_branch(ctx *Simple_hierarchical_branchContext) {
}

// ExitSimple_hierarchical_branch is called when production simple_hierarchical_branch is exited.
func (s *BaseVerilogListener) ExitSimple_hierarchical_branch(ctx *Simple_hierarchical_branchContext) {
}

// EnterEscaped_hierarchical_branch is called when production escaped_hierarchical_branch is entered.
func (s *BaseVerilogListener) EnterEscaped_hierarchical_branch(ctx *Escaped_hierarchical_branchContext) {
}

// ExitEscaped_hierarchical_branch is called when production escaped_hierarchical_branch is exited.
func (s *BaseVerilogListener) ExitEscaped_hierarchical_branch(ctx *Escaped_hierarchical_branchContext) {
}
