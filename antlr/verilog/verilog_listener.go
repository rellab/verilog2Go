// Code generated from Verilog.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Verilog

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VerilogListener is a complete listener for a parse tree produced by VerilogParser.
type VerilogListener interface {
	antlr.ParseTreeListener

	// EnterConfig_declaration is called when entering the config_declaration production.
	EnterConfig_declaration(c *Config_declarationContext)

	// EnterDesign_statement is called when entering the design_statement production.
	EnterDesign_statement(c *Design_statementContext)

	// EnterConfig_rule_statement is called when entering the config_rule_statement production.
	EnterConfig_rule_statement(c *Config_rule_statementContext)

	// EnterDefault_clause is called when entering the default_clause production.
	EnterDefault_clause(c *Default_clauseContext)

	// EnterInst_clause is called when entering the inst_clause production.
	EnterInst_clause(c *Inst_clauseContext)

	// EnterInst_name is called when entering the inst_name production.
	EnterInst_name(c *Inst_nameContext)

	// EnterLiblist_clause is called when entering the liblist_clause production.
	EnterLiblist_clause(c *Liblist_clauseContext)

	// EnterCell_clause is called when entering the cell_clause production.
	EnterCell_clause(c *Cell_clauseContext)

	// EnterUse_clause is called when entering the use_clause production.
	EnterUse_clause(c *Use_clauseContext)

	// EnterSource_text is called when entering the source_text production.
	EnterSource_text(c *Source_textContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// EnterModule_declaration is called when entering the module_declaration production.
	EnterModule_declaration(c *Module_declarationContext)

	// EnterModule_keyword is called when entering the module_keyword production.
	EnterModule_keyword(c *Module_keywordContext)

	// EnterModule_parameter_port_list is called when entering the module_parameter_port_list production.
	EnterModule_parameter_port_list(c *Module_parameter_port_listContext)

	// EnterList_of_ports is called when entering the list_of_ports production.
	EnterList_of_ports(c *List_of_portsContext)

	// EnterList_of_port_declarations is called when entering the list_of_port_declarations production.
	EnterList_of_port_declarations(c *List_of_port_declarationsContext)

	// EnterPort is called when entering the port production.
	EnterPort(c *PortContext)

	// EnterPort_expression is called when entering the port_expression production.
	EnterPort_expression(c *Port_expressionContext)

	// EnterPort_reference is called when entering the port_reference production.
	EnterPort_reference(c *Port_referenceContext)

	// EnterPort_declaration is called when entering the port_declaration production.
	EnterPort_declaration(c *Port_declarationContext)

	// EnterModule_item is called when entering the module_item production.
	EnterModule_item(c *Module_itemContext)

	// EnterModule_or_generate_item is called when entering the module_or_generate_item production.
	EnterModule_or_generate_item(c *Module_or_generate_itemContext)

	// EnterNon_port_module_item is called when entering the non_port_module_item production.
	EnterNon_port_module_item(c *Non_port_module_itemContext)

	// EnterModule_or_generate_item_declaration is called when entering the module_or_generate_item_declaration production.
	EnterModule_or_generate_item_declaration(c *Module_or_generate_item_declarationContext)

	// EnterParameter_override is called when entering the parameter_override production.
	EnterParameter_override(c *Parameter_overrideContext)

	// EnterLocal_parameter_declaration is called when entering the local_parameter_declaration production.
	EnterLocal_parameter_declaration(c *Local_parameter_declarationContext)

	// EnterParameter_declaration is called when entering the parameter_declaration production.
	EnterParameter_declaration(c *Parameter_declarationContext)

	// EnterParameter_declaration_ is called when entering the parameter_declaration_ production.
	EnterParameter_declaration_(c *Parameter_declaration_Context)

	// EnterSpecparam_declaration is called when entering the specparam_declaration production.
	EnterSpecparam_declaration(c *Specparam_declarationContext)

	// EnterInout_declaration is called when entering the inout_declaration production.
	EnterInout_declaration(c *Inout_declarationContext)

	// EnterInput_declaration is called when entering the input_declaration production.
	EnterInput_declaration(c *Input_declarationContext)

	// EnterOutput_declaration is called when entering the output_declaration production.
	EnterOutput_declaration(c *Output_declarationContext)

	// EnterEvent_declaration is called when entering the event_declaration production.
	EnterEvent_declaration(c *Event_declarationContext)

	// EnterGenvar_declaration is called when entering the genvar_declaration production.
	EnterGenvar_declaration(c *Genvar_declarationContext)

	// EnterInteger_declaration is called when entering the integer_declaration production.
	EnterInteger_declaration(c *Integer_declarationContext)

	// EnterTime_declaration is called when entering the time_declaration production.
	EnterTime_declaration(c *Time_declarationContext)

	// EnterReal_declaration is called when entering the real_declaration production.
	EnterReal_declaration(c *Real_declarationContext)

	// EnterRealtime_declaration is called when entering the realtime_declaration production.
	EnterRealtime_declaration(c *Realtime_declarationContext)

	// EnterReg_declaration is called when entering the reg_declaration production.
	EnterReg_declaration(c *Reg_declarationContext)

	// EnterNet_declaration is called when entering the net_declaration production.
	EnterNet_declaration(c *Net_declarationContext)

	// EnterNet_type is called when entering the net_type production.
	EnterNet_type(c *Net_typeContext)

	// EnterOutput_variable_type is called when entering the output_variable_type production.
	EnterOutput_variable_type(c *Output_variable_typeContext)

	// EnterReal_type is called when entering the real_type production.
	EnterReal_type(c *Real_typeContext)

	// EnterVariable_type is called when entering the variable_type production.
	EnterVariable_type(c *Variable_typeContext)

	// EnterDrive_strength is called when entering the drive_strength production.
	EnterDrive_strength(c *Drive_strengthContext)

	// EnterStrength0 is called when entering the strength0 production.
	EnterStrength0(c *Strength0Context)

	// EnterStrength1 is called when entering the strength1 production.
	EnterStrength1(c *Strength1Context)

	// EnterCharge_strength is called when entering the charge_strength production.
	EnterCharge_strength(c *Charge_strengthContext)

	// EnterDelay3 is called when entering the delay3 production.
	EnterDelay3(c *Delay3Context)

	// EnterDelay2 is called when entering the delay2 production.
	EnterDelay2(c *Delay2Context)

	// EnterDelay_value is called when entering the delay_value production.
	EnterDelay_value(c *Delay_valueContext)

	// EnterList_of_event_identifiers is called when entering the list_of_event_identifiers production.
	EnterList_of_event_identifiers(c *List_of_event_identifiersContext)

	// EnterList_of_net_identifiers is called when entering the list_of_net_identifiers production.
	EnterList_of_net_identifiers(c *List_of_net_identifiersContext)

	// EnterList_of_genvar_identifiers is called when entering the list_of_genvar_identifiers production.
	EnterList_of_genvar_identifiers(c *List_of_genvar_identifiersContext)

	// EnterList_of_port_identifiers is called when entering the list_of_port_identifiers production.
	EnterList_of_port_identifiers(c *List_of_port_identifiersContext)

	// EnterList_of_net_decl_assignments is called when entering the list_of_net_decl_assignments production.
	EnterList_of_net_decl_assignments(c *List_of_net_decl_assignmentsContext)

	// EnterList_of_param_assignments is called when entering the list_of_param_assignments production.
	EnterList_of_param_assignments(c *List_of_param_assignmentsContext)

	// EnterList_of_specparam_assignments is called when entering the list_of_specparam_assignments production.
	EnterList_of_specparam_assignments(c *List_of_specparam_assignmentsContext)

	// EnterList_of_real_identifiers is called when entering the list_of_real_identifiers production.
	EnterList_of_real_identifiers(c *List_of_real_identifiersContext)

	// EnterList_of_variable_identifiers is called when entering the list_of_variable_identifiers production.
	EnterList_of_variable_identifiers(c *List_of_variable_identifiersContext)

	// EnterList_of_variable_port_identifiers is called when entering the list_of_variable_port_identifiers production.
	EnterList_of_variable_port_identifiers(c *List_of_variable_port_identifiersContext)

	// EnterNet_decl_assignment is called when entering the net_decl_assignment production.
	EnterNet_decl_assignment(c *Net_decl_assignmentContext)

	// EnterParam_assignment is called when entering the param_assignment production.
	EnterParam_assignment(c *Param_assignmentContext)

	// EnterSpecparam_assignment is called when entering the specparam_assignment production.
	EnterSpecparam_assignment(c *Specparam_assignmentContext)

	// EnterPulse_control_specparam is called when entering the pulse_control_specparam production.
	EnterPulse_control_specparam(c *Pulse_control_specparamContext)

	// EnterError_limit_value is called when entering the error_limit_value production.
	EnterError_limit_value(c *Error_limit_valueContext)

	// EnterReject_limit_value is called when entering the reject_limit_value production.
	EnterReject_limit_value(c *Reject_limit_valueContext)

	// EnterLimit_value is called when entering the limit_value production.
	EnterLimit_value(c *Limit_valueContext)

	// EnterDimension is called when entering the dimension production.
	EnterDimension(c *DimensionContext)

	// EnterRange_ is called when entering the range_ production.
	EnterRange_(c *Range_Context)

	// EnterFunction_declaration is called when entering the function_declaration production.
	EnterFunction_declaration(c *Function_declarationContext)

	// EnterFunction_item_declaration is called when entering the function_item_declaration production.
	EnterFunction_item_declaration(c *Function_item_declarationContext)

	// EnterFunction_port_list is called when entering the function_port_list production.
	EnterFunction_port_list(c *Function_port_listContext)

	// EnterFunction_port is called when entering the function_port production.
	EnterFunction_port(c *Function_portContext)

	// EnterRange_or_type is called when entering the range_or_type production.
	EnterRange_or_type(c *Range_or_typeContext)

	// EnterTask_declaration is called when entering the task_declaration production.
	EnterTask_declaration(c *Task_declarationContext)

	// EnterTask_item_declaration is called when entering the task_item_declaration production.
	EnterTask_item_declaration(c *Task_item_declarationContext)

	// EnterTask_port_list is called when entering the task_port_list production.
	EnterTask_port_list(c *Task_port_listContext)

	// EnterTask_port_item is called when entering the task_port_item production.
	EnterTask_port_item(c *Task_port_itemContext)

	// EnterTf_decl_header is called when entering the tf_decl_header production.
	EnterTf_decl_header(c *Tf_decl_headerContext)

	// EnterTf_declaration is called when entering the tf_declaration production.
	EnterTf_declaration(c *Tf_declarationContext)

	// EnterTask_port_type is called when entering the task_port_type production.
	EnterTask_port_type(c *Task_port_typeContext)

	// EnterBlock_item_declaration is called when entering the block_item_declaration production.
	EnterBlock_item_declaration(c *Block_item_declarationContext)

	// EnterBlock_reg_declaration is called when entering the block_reg_declaration production.
	EnterBlock_reg_declaration(c *Block_reg_declarationContext)

	// EnterList_of_block_variable_identifiers is called when entering the list_of_block_variable_identifiers production.
	EnterList_of_block_variable_identifiers(c *List_of_block_variable_identifiersContext)

	// EnterBlock_variable_type is called when entering the block_variable_type production.
	EnterBlock_variable_type(c *Block_variable_typeContext)

	// EnterGate_instantiation is called when entering the gate_instantiation production.
	EnterGate_instantiation(c *Gate_instantiationContext)

	// EnterCmos_switch_instance is called when entering the cmos_switch_instance production.
	EnterCmos_switch_instance(c *Cmos_switch_instanceContext)

	// EnterEnable_gate_instance is called when entering the enable_gate_instance production.
	EnterEnable_gate_instance(c *Enable_gate_instanceContext)

	// EnterMos_switch_instance is called when entering the mos_switch_instance production.
	EnterMos_switch_instance(c *Mos_switch_instanceContext)

	// EnterN_input_gate_instance is called when entering the n_input_gate_instance production.
	EnterN_input_gate_instance(c *N_input_gate_instanceContext)

	// EnterN_output_gate_instance is called when entering the n_output_gate_instance production.
	EnterN_output_gate_instance(c *N_output_gate_instanceContext)

	// EnterPass_switch_instance is called when entering the pass_switch_instance production.
	EnterPass_switch_instance(c *Pass_switch_instanceContext)

	// EnterPass_enable_switch_instance is called when entering the pass_enable_switch_instance production.
	EnterPass_enable_switch_instance(c *Pass_enable_switch_instanceContext)

	// EnterPull_gate_instance is called when entering the pull_gate_instance production.
	EnterPull_gate_instance(c *Pull_gate_instanceContext)

	// EnterName_of_gate_instance is called when entering the name_of_gate_instance production.
	EnterName_of_gate_instance(c *Name_of_gate_instanceContext)

	// EnterPulldown_strength is called when entering the pulldown_strength production.
	EnterPulldown_strength(c *Pulldown_strengthContext)

	// EnterPullup_strength is called when entering the pullup_strength production.
	EnterPullup_strength(c *Pullup_strengthContext)

	// EnterEnable_terminal is called when entering the enable_terminal production.
	EnterEnable_terminal(c *Enable_terminalContext)

	// EnterNcontrol_terminal is called when entering the ncontrol_terminal production.
	EnterNcontrol_terminal(c *Ncontrol_terminalContext)

	// EnterPcontrol_terminal is called when entering the pcontrol_terminal production.
	EnterPcontrol_terminal(c *Pcontrol_terminalContext)

	// EnterInput_terminal is called when entering the input_terminal production.
	EnterInput_terminal(c *Input_terminalContext)

	// EnterInout_terminal is called when entering the inout_terminal production.
	EnterInout_terminal(c *Inout_terminalContext)

	// EnterOutput_terminal is called when entering the output_terminal production.
	EnterOutput_terminal(c *Output_terminalContext)

	// EnterCmos_switchtype is called when entering the cmos_switchtype production.
	EnterCmos_switchtype(c *Cmos_switchtypeContext)

	// EnterEnable_gatetype is called when entering the enable_gatetype production.
	EnterEnable_gatetype(c *Enable_gatetypeContext)

	// EnterMos_switchtype is called when entering the mos_switchtype production.
	EnterMos_switchtype(c *Mos_switchtypeContext)

	// EnterN_input_gatetype is called when entering the n_input_gatetype production.
	EnterN_input_gatetype(c *N_input_gatetypeContext)

	// EnterN_output_gatetype is called when entering the n_output_gatetype production.
	EnterN_output_gatetype(c *N_output_gatetypeContext)

	// EnterPass_en_switchtype is called when entering the pass_en_switchtype production.
	EnterPass_en_switchtype(c *Pass_en_switchtypeContext)

	// EnterPass_switchtype is called when entering the pass_switchtype production.
	EnterPass_switchtype(c *Pass_switchtypeContext)

	// EnterModule_instantiation is called when entering the module_instantiation production.
	EnterModule_instantiation(c *Module_instantiationContext)

	// EnterParameter_value_assignment is called when entering the parameter_value_assignment production.
	EnterParameter_value_assignment(c *Parameter_value_assignmentContext)

	// EnterList_of_parameter_assignments is called when entering the list_of_parameter_assignments production.
	EnterList_of_parameter_assignments(c *List_of_parameter_assignmentsContext)

	// EnterOrdered_parameter_assignment is called when entering the ordered_parameter_assignment production.
	EnterOrdered_parameter_assignment(c *Ordered_parameter_assignmentContext)

	// EnterNamed_parameter_assignment is called when entering the named_parameter_assignment production.
	EnterNamed_parameter_assignment(c *Named_parameter_assignmentContext)

	// EnterModule_instance is called when entering the module_instance production.
	EnterModule_instance(c *Module_instanceContext)

	// EnterName_of_instance is called when entering the name_of_instance production.
	EnterName_of_instance(c *Name_of_instanceContext)

	// EnterList_of_port_connections is called when entering the list_of_port_connections production.
	EnterList_of_port_connections(c *List_of_port_connectionsContext)

	// EnterOrdered_port_connection is called when entering the ordered_port_connection production.
	EnterOrdered_port_connection(c *Ordered_port_connectionContext)

	// EnterNamed_port_connection is called when entering the named_port_connection production.
	EnterNamed_port_connection(c *Named_port_connectionContext)

	// EnterGenerated_instantiation is called when entering the generated_instantiation production.
	EnterGenerated_instantiation(c *Generated_instantiationContext)

	// EnterGenerate_item_or_null is called when entering the generate_item_or_null production.
	EnterGenerate_item_or_null(c *Generate_item_or_nullContext)

	// EnterGenerate_item is called when entering the generate_item production.
	EnterGenerate_item(c *Generate_itemContext)

	// EnterGenerate_conditional_statement is called when entering the generate_conditional_statement production.
	EnterGenerate_conditional_statement(c *Generate_conditional_statementContext)

	// EnterGenerate_case_statement is called when entering the generate_case_statement production.
	EnterGenerate_case_statement(c *Generate_case_statementContext)

	// EnterGenvar_case_item is called when entering the genvar_case_item production.
	EnterGenvar_case_item(c *Genvar_case_itemContext)

	// EnterGenerate_loop_statement is called when entering the generate_loop_statement production.
	EnterGenerate_loop_statement(c *Generate_loop_statementContext)

	// EnterGenvar_assignment is called when entering the genvar_assignment production.
	EnterGenvar_assignment(c *Genvar_assignmentContext)

	// EnterGenerate_block is called when entering the generate_block production.
	EnterGenerate_block(c *Generate_blockContext)

	// EnterContinuous_assign is called when entering the continuous_assign production.
	EnterContinuous_assign(c *Continuous_assignContext)

	// EnterList_of_net_assignments is called when entering the list_of_net_assignments production.
	EnterList_of_net_assignments(c *List_of_net_assignmentsContext)

	// EnterNet_assignment is called when entering the net_assignment production.
	EnterNet_assignment(c *Net_assignmentContext)

	// EnterInitial_construct is called when entering the initial_construct production.
	EnterInitial_construct(c *Initial_constructContext)

	// EnterAlways_construct is called when entering the always_construct production.
	EnterAlways_construct(c *Always_constructContext)

	// EnterBlocking_assignment is called when entering the blocking_assignment production.
	EnterBlocking_assignment(c *Blocking_assignmentContext)

	// EnterNonblocking_assignment is called when entering the nonblocking_assignment production.
	EnterNonblocking_assignment(c *Nonblocking_assignmentContext)

	// EnterProcedural_continuous_assignments is called when entering the procedural_continuous_assignments production.
	EnterProcedural_continuous_assignments(c *Procedural_continuous_assignmentsContext)

	// EnterFunction_blocking_assignment is called when entering the function_blocking_assignment production.
	EnterFunction_blocking_assignment(c *Function_blocking_assignmentContext)

	// EnterFunction_statement_or_null is called when entering the function_statement_or_null production.
	EnterFunction_statement_or_null(c *Function_statement_or_nullContext)

	// EnterFunction_seq_block is called when entering the function_seq_block production.
	EnterFunction_seq_block(c *Function_seq_blockContext)

	// EnterVariable_assignment is called when entering the variable_assignment production.
	EnterVariable_assignment(c *Variable_assignmentContext)

	// EnterPar_block is called when entering the par_block production.
	EnterPar_block(c *Par_blockContext)

	// EnterSeq_block is called when entering the seq_block production.
	EnterSeq_block(c *Seq_blockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterStatement_or_null is called when entering the statement_or_null production.
	EnterStatement_or_null(c *Statement_or_nullContext)

	// EnterFunction_statement is called when entering the function_statement production.
	EnterFunction_statement(c *Function_statementContext)

	// EnterDelay_or_event_control is called when entering the delay_or_event_control production.
	EnterDelay_or_event_control(c *Delay_or_event_controlContext)

	// EnterDelay_control is called when entering the delay_control production.
	EnterDelay_control(c *Delay_controlContext)

	// EnterDisable_statement is called when entering the disable_statement production.
	EnterDisable_statement(c *Disable_statementContext)

	// EnterEvent_control is called when entering the event_control production.
	EnterEvent_control(c *Event_controlContext)

	// EnterEvent_trigger is called when entering the event_trigger production.
	EnterEvent_trigger(c *Event_triggerContext)

	// EnterEvent_expression is called when entering the event_expression production.
	EnterEvent_expression(c *Event_expressionContext)

	// EnterEvent_primary is called when entering the event_primary production.
	EnterEvent_primary(c *Event_primaryContext)

	// EnterProcedural_timing_control_statement is called when entering the procedural_timing_control_statement production.
	EnterProcedural_timing_control_statement(c *Procedural_timing_control_statementContext)

	// EnterWait_statement is called when entering the wait_statement production.
	EnterWait_statement(c *Wait_statementContext)

	// EnterConditional_statement is called when entering the conditional_statement production.
	EnterConditional_statement(c *Conditional_statementContext)

	// EnterIf_else_if_statement is called when entering the if_else_if_statement production.
	EnterIf_else_if_statement(c *If_else_if_statementContext)

	// EnterElse_statement is called when entering the else_statement production.
	EnterElse_statement(c *Else_statementContext)

	// EnterFunction_conditional_statement is called when entering the function_conditional_statement production.
	EnterFunction_conditional_statement(c *Function_conditional_statementContext)

	// EnterFunction_if_else_if_statement is called when entering the function_if_else_if_statement production.
	EnterFunction_if_else_if_statement(c *Function_if_else_if_statementContext)

	// EnterCase_statement is called when entering the case_statement production.
	EnterCase_statement(c *Case_statementContext)

	// EnterCase_item is called when entering the case_item production.
	EnterCase_item(c *Case_itemContext)

	// EnterFunction_case_statement is called when entering the function_case_statement production.
	EnterFunction_case_statement(c *Function_case_statementContext)

	// EnterFunction_case_item is called when entering the function_case_item production.
	EnterFunction_case_item(c *Function_case_itemContext)

	// EnterFunction_loop_statement is called when entering the function_loop_statement production.
	EnterFunction_loop_statement(c *Function_loop_statementContext)

	// EnterLoop_statement is called when entering the loop_statement production.
	EnterLoop_statement(c *Loop_statementContext)

	// EnterSystem_task_enable is called when entering the system_task_enable production.
	EnterSystem_task_enable(c *System_task_enableContext)

	// EnterTask_enable is called when entering the task_enable production.
	EnterTask_enable(c *Task_enableContext)

	// EnterSpecify_block is called when entering the specify_block production.
	EnterSpecify_block(c *Specify_blockContext)

	// EnterSpecify_item is called when entering the specify_item production.
	EnterSpecify_item(c *Specify_itemContext)

	// EnterPulsestyle_declaration is called when entering the pulsestyle_declaration production.
	EnterPulsestyle_declaration(c *Pulsestyle_declarationContext)

	// EnterShowcancelled_declaration is called when entering the showcancelled_declaration production.
	EnterShowcancelled_declaration(c *Showcancelled_declarationContext)

	// EnterPath_declaration is called when entering the path_declaration production.
	EnterPath_declaration(c *Path_declarationContext)

	// EnterSimple_path_declaration is called when entering the simple_path_declaration production.
	EnterSimple_path_declaration(c *Simple_path_declarationContext)

	// EnterParallel_path_description is called when entering the parallel_path_description production.
	EnterParallel_path_description(c *Parallel_path_descriptionContext)

	// EnterFull_path_description is called when entering the full_path_description production.
	EnterFull_path_description(c *Full_path_descriptionContext)

	// EnterList_of_path_inputs is called when entering the list_of_path_inputs production.
	EnterList_of_path_inputs(c *List_of_path_inputsContext)

	// EnterList_of_path_outputs is called when entering the list_of_path_outputs production.
	EnterList_of_path_outputs(c *List_of_path_outputsContext)

	// EnterSpecify_input_terminal_descriptor is called when entering the specify_input_terminal_descriptor production.
	EnterSpecify_input_terminal_descriptor(c *Specify_input_terminal_descriptorContext)

	// EnterSpecify_output_terminal_descriptor is called when entering the specify_output_terminal_descriptor production.
	EnterSpecify_output_terminal_descriptor(c *Specify_output_terminal_descriptorContext)

	// EnterInput_identifier is called when entering the input_identifier production.
	EnterInput_identifier(c *Input_identifierContext)

	// EnterOutput_identifier is called when entering the output_identifier production.
	EnterOutput_identifier(c *Output_identifierContext)

	// EnterPath_delay_value is called when entering the path_delay_value production.
	EnterPath_delay_value(c *Path_delay_valueContext)

	// EnterList_of_path_delay_expressions is called when entering the list_of_path_delay_expressions production.
	EnterList_of_path_delay_expressions(c *List_of_path_delay_expressionsContext)

	// EnterT_path_delay_expression is called when entering the t_path_delay_expression production.
	EnterT_path_delay_expression(c *T_path_delay_expressionContext)

	// EnterTrise_path_delay_expression is called when entering the trise_path_delay_expression production.
	EnterTrise_path_delay_expression(c *Trise_path_delay_expressionContext)

	// EnterTfall_path_delay_expression is called when entering the tfall_path_delay_expression production.
	EnterTfall_path_delay_expression(c *Tfall_path_delay_expressionContext)

	// EnterTz_path_delay_expression is called when entering the tz_path_delay_expression production.
	EnterTz_path_delay_expression(c *Tz_path_delay_expressionContext)

	// EnterT01_path_delay_expression is called when entering the t01_path_delay_expression production.
	EnterT01_path_delay_expression(c *T01_path_delay_expressionContext)

	// EnterT10_path_delay_expression is called when entering the t10_path_delay_expression production.
	EnterT10_path_delay_expression(c *T10_path_delay_expressionContext)

	// EnterT0z_path_delay_expression is called when entering the t0z_path_delay_expression production.
	EnterT0z_path_delay_expression(c *T0z_path_delay_expressionContext)

	// EnterTz1_path_delay_expression is called when entering the tz1_path_delay_expression production.
	EnterTz1_path_delay_expression(c *Tz1_path_delay_expressionContext)

	// EnterT1z_path_delay_expression is called when entering the t1z_path_delay_expression production.
	EnterT1z_path_delay_expression(c *T1z_path_delay_expressionContext)

	// EnterTz0_path_delay_expression is called when entering the tz0_path_delay_expression production.
	EnterTz0_path_delay_expression(c *Tz0_path_delay_expressionContext)

	// EnterT0x_path_delay_expression is called when entering the t0x_path_delay_expression production.
	EnterT0x_path_delay_expression(c *T0x_path_delay_expressionContext)

	// EnterTx1_path_delay_expression is called when entering the tx1_path_delay_expression production.
	EnterTx1_path_delay_expression(c *Tx1_path_delay_expressionContext)

	// EnterT1x_path_delay_expression is called when entering the t1x_path_delay_expression production.
	EnterT1x_path_delay_expression(c *T1x_path_delay_expressionContext)

	// EnterTx0_path_delay_expression is called when entering the tx0_path_delay_expression production.
	EnterTx0_path_delay_expression(c *Tx0_path_delay_expressionContext)

	// EnterTxz_path_delay_expression is called when entering the txz_path_delay_expression production.
	EnterTxz_path_delay_expression(c *Txz_path_delay_expressionContext)

	// EnterTzx_path_delay_expression is called when entering the tzx_path_delay_expression production.
	EnterTzx_path_delay_expression(c *Tzx_path_delay_expressionContext)

	// EnterPath_delay_expression is called when entering the path_delay_expression production.
	EnterPath_delay_expression(c *Path_delay_expressionContext)

	// EnterEdge_sensitive_path_declaration is called when entering the edge_sensitive_path_declaration production.
	EnterEdge_sensitive_path_declaration(c *Edge_sensitive_path_declarationContext)

	// EnterParallel_edge_sensitive_path_description is called when entering the parallel_edge_sensitive_path_description production.
	EnterParallel_edge_sensitive_path_description(c *Parallel_edge_sensitive_path_descriptionContext)

	// EnterFull_edge_sensitive_path_description is called when entering the full_edge_sensitive_path_description production.
	EnterFull_edge_sensitive_path_description(c *Full_edge_sensitive_path_descriptionContext)

	// EnterData_source_expression is called when entering the data_source_expression production.
	EnterData_source_expression(c *Data_source_expressionContext)

	// EnterEdge_identifier is called when entering the edge_identifier production.
	EnterEdge_identifier(c *Edge_identifierContext)

	// EnterState_dependent_path_declaration is called when entering the state_dependent_path_declaration production.
	EnterState_dependent_path_declaration(c *State_dependent_path_declarationContext)

	// EnterPolarity_operator is called when entering the polarity_operator production.
	EnterPolarity_operator(c *Polarity_operatorContext)

	// EnterChecktime_condition is called when entering the checktime_condition production.
	EnterChecktime_condition(c *Checktime_conditionContext)

	// EnterDelayed_data is called when entering the delayed_data production.
	EnterDelayed_data(c *Delayed_dataContext)

	// EnterDelayed_reference is called when entering the delayed_reference production.
	EnterDelayed_reference(c *Delayed_referenceContext)

	// EnterEnd_edge_offset is called when entering the end_edge_offset production.
	EnterEnd_edge_offset(c *End_edge_offsetContext)

	// EnterEvent_based_flag is called when entering the event_based_flag production.
	EnterEvent_based_flag(c *Event_based_flagContext)

	// EnterNotify_reg is called when entering the notify_reg production.
	EnterNotify_reg(c *Notify_regContext)

	// EnterRemain_active_flag is called when entering the remain_active_flag production.
	EnterRemain_active_flag(c *Remain_active_flagContext)

	// EnterStamptime_condition is called when entering the stamptime_condition production.
	EnterStamptime_condition(c *Stamptime_conditionContext)

	// EnterStart_edge_offset is called when entering the start_edge_offset production.
	EnterStart_edge_offset(c *Start_edge_offsetContext)

	// EnterThreshold is called when entering the threshold production.
	EnterThreshold(c *ThresholdContext)

	// EnterTiming_check_limit is called when entering the timing_check_limit production.
	EnterTiming_check_limit(c *Timing_check_limitContext)

	// EnterConcatenation is called when entering the concatenation production.
	EnterConcatenation(c *ConcatenationContext)

	// EnterConstant_concatenation is called when entering the constant_concatenation production.
	EnterConstant_concatenation(c *Constant_concatenationContext)

	// EnterConstant_multiple_concatenation is called when entering the constant_multiple_concatenation production.
	EnterConstant_multiple_concatenation(c *Constant_multiple_concatenationContext)

	// EnterModule_path_concatenation is called when entering the module_path_concatenation production.
	EnterModule_path_concatenation(c *Module_path_concatenationContext)

	// EnterModule_path_multiple_concatenation is called when entering the module_path_multiple_concatenation production.
	EnterModule_path_multiple_concatenation(c *Module_path_multiple_concatenationContext)

	// EnterMultiple_concatenation is called when entering the multiple_concatenation production.
	EnterMultiple_concatenation(c *Multiple_concatenationContext)

	// EnterNet_concatenation is called when entering the net_concatenation production.
	EnterNet_concatenation(c *Net_concatenationContext)

	// EnterNet_concatenation_value is called when entering the net_concatenation_value production.
	EnterNet_concatenation_value(c *Net_concatenation_valueContext)

	// EnterVariable_concatenation is called when entering the variable_concatenation production.
	EnterVariable_concatenation(c *Variable_concatenationContext)

	// EnterVariable_concatenation_value is called when entering the variable_concatenation_value production.
	EnterVariable_concatenation_value(c *Variable_concatenation_valueContext)

	// EnterConstant_function_call is called when entering the constant_function_call production.
	EnterConstant_function_call(c *Constant_function_callContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterSystem_function_call is called when entering the system_function_call production.
	EnterSystem_function_call(c *System_function_callContext)

	// EnterGenvar_function_call is called when entering the genvar_function_call production.
	EnterGenvar_function_call(c *Genvar_function_callContext)

	// EnterBase_expression is called when entering the base_expression production.
	EnterBase_expression(c *Base_expressionContext)

	// EnterConstant_base_expression is called when entering the constant_base_expression production.
	EnterConstant_base_expression(c *Constant_base_expressionContext)

	// EnterConstant_expression is called when entering the constant_expression production.
	EnterConstant_expression(c *Constant_expressionContext)

	// EnterConstant_mintypmax_expression is called when entering the constant_mintypmax_expression production.
	EnterConstant_mintypmax_expression(c *Constant_mintypmax_expressionContext)

	// EnterConstant_range_expression is called when entering the constant_range_expression production.
	EnterConstant_range_expression(c *Constant_range_expressionContext)

	// EnterDimension_constant_expression is called when entering the dimension_constant_expression production.
	EnterDimension_constant_expression(c *Dimension_constant_expressionContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterLsb_constant_expression is called when entering the lsb_constant_expression production.
	EnterLsb_constant_expression(c *Lsb_constant_expressionContext)

	// EnterMintypmax_expression is called when entering the mintypmax_expression production.
	EnterMintypmax_expression(c *Mintypmax_expressionContext)

	// EnterModule_path_conditional_expression is called when entering the module_path_conditional_expression production.
	EnterModule_path_conditional_expression(c *Module_path_conditional_expressionContext)

	// EnterModule_path_expression is called when entering the module_path_expression production.
	EnterModule_path_expression(c *Module_path_expressionContext)

	// EnterModule_path_mintypmax_expression is called when entering the module_path_mintypmax_expression production.
	EnterModule_path_mintypmax_expression(c *Module_path_mintypmax_expressionContext)

	// EnterMsb_constant_expression is called when entering the msb_constant_expression production.
	EnterMsb_constant_expression(c *Msb_constant_expressionContext)

	// EnterRange_expression is called when entering the range_expression production.
	EnterRange_expression(c *Range_expressionContext)

	// EnterWidth_constant_expression is called when entering the width_constant_expression production.
	EnterWidth_constant_expression(c *Width_constant_expressionContext)

	// EnterConstant_primary is called when entering the constant_primary production.
	EnterConstant_primary(c *Constant_primaryContext)

	// EnterModule_path_primary is called when entering the module_path_primary production.
	EnterModule_path_primary(c *Module_path_primaryContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterNet_lvalue is called when entering the net_lvalue production.
	EnterNet_lvalue(c *Net_lvalueContext)

	// EnterVariable_lvalue is called when entering the variable_lvalue production.
	EnterVariable_lvalue(c *Variable_lvalueContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterBinary_operator is called when entering the binary_operator production.
	EnterBinary_operator(c *Binary_operatorContext)

	// EnterUnary_module_path_operator is called when entering the unary_module_path_operator production.
	EnterUnary_module_path_operator(c *Unary_module_path_operatorContext)

	// EnterBinary_module_path_operator is called when entering the binary_module_path_operator production.
	EnterBinary_module_path_operator(c *Binary_module_path_operatorContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterTiming_spec is called when entering the timing_spec production.
	EnterTiming_spec(c *Timing_specContext)

	// EnterDefine is called when entering the define production.
	EnterDefine(c *DefineContext)

	// EnterAttribute_instance is called when entering the attribute_instance production.
	EnterAttribute_instance(c *Attribute_instanceContext)

	// EnterAttr_spec is called when entering the attr_spec production.
	EnterAttr_spec(c *Attr_specContext)

	// EnterAttr_name is called when entering the attr_name production.
	EnterAttr_name(c *Attr_nameContext)

	// EnterArrayed_identifier is called when entering the arrayed_identifier production.
	EnterArrayed_identifier(c *Arrayed_identifierContext)

	// EnterBlock_identifier is called when entering the block_identifier production.
	EnterBlock_identifier(c *Block_identifierContext)

	// EnterCell_identifier is called when entering the cell_identifier production.
	EnterCell_identifier(c *Cell_identifierContext)

	// EnterConfig_identifier is called when entering the config_identifier production.
	EnterConfig_identifier(c *Config_identifierContext)

	// EnterDefine_identifier is called when entering the define_identifier production.
	EnterDefine_identifier(c *Define_identifierContext)

	// EnterEscaped_arrayed_identifier is called when entering the escaped_arrayed_identifier production.
	EnterEscaped_arrayed_identifier(c *Escaped_arrayed_identifierContext)

	// EnterEscaped_hierarchical_identifier is called when entering the escaped_hierarchical_identifier production.
	EnterEscaped_hierarchical_identifier(c *Escaped_hierarchical_identifierContext)

	// EnterEvent_identifier is called when entering the event_identifier production.
	EnterEvent_identifier(c *Event_identifierContext)

	// EnterFunction_identifier is called when entering the function_identifier production.
	EnterFunction_identifier(c *Function_identifierContext)

	// EnterGate_instance_identifier is called when entering the gate_instance_identifier production.
	EnterGate_instance_identifier(c *Gate_instance_identifierContext)

	// EnterGenerate_block_identifier is called when entering the generate_block_identifier production.
	EnterGenerate_block_identifier(c *Generate_block_identifierContext)

	// EnterGenvar_function_identifier is called when entering the genvar_function_identifier production.
	EnterGenvar_function_identifier(c *Genvar_function_identifierContext)

	// EnterGenvar_identifier is called when entering the genvar_identifier production.
	EnterGenvar_identifier(c *Genvar_identifierContext)

	// EnterHierarchical_block_identifier is called when entering the hierarchical_block_identifier production.
	EnterHierarchical_block_identifier(c *Hierarchical_block_identifierContext)

	// EnterHierarchical_event_identifier is called when entering the hierarchical_event_identifier production.
	EnterHierarchical_event_identifier(c *Hierarchical_event_identifierContext)

	// EnterHierarchical_function_identifier is called when entering the hierarchical_function_identifier production.
	EnterHierarchical_function_identifier(c *Hierarchical_function_identifierContext)

	// EnterHierarchical_identifier is called when entering the hierarchical_identifier production.
	EnterHierarchical_identifier(c *Hierarchical_identifierContext)

	// EnterHierarchical_net_identifier is called when entering the hierarchical_net_identifier production.
	EnterHierarchical_net_identifier(c *Hierarchical_net_identifierContext)

	// EnterHierarchical_variable_identifier is called when entering the hierarchical_variable_identifier production.
	EnterHierarchical_variable_identifier(c *Hierarchical_variable_identifierContext)

	// EnterHierarchical_task_identifier is called when entering the hierarchical_task_identifier production.
	EnterHierarchical_task_identifier(c *Hierarchical_task_identifierContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterInout_port_identifier is called when entering the inout_port_identifier production.
	EnterInout_port_identifier(c *Inout_port_identifierContext)

	// EnterInput_port_identifier is called when entering the input_port_identifier production.
	EnterInput_port_identifier(c *Input_port_identifierContext)

	// EnterInstance_identifier is called when entering the instance_identifier production.
	EnterInstance_identifier(c *Instance_identifierContext)

	// EnterLibrary_identifier is called when entering the library_identifier production.
	EnterLibrary_identifier(c *Library_identifierContext)

	// EnterMemory_identifier is called when entering the memory_identifier production.
	EnterMemory_identifier(c *Memory_identifierContext)

	// EnterModule_identifier is called when entering the module_identifier production.
	EnterModule_identifier(c *Module_identifierContext)

	// EnterModule_instance_identifier is called when entering the module_instance_identifier production.
	EnterModule_instance_identifier(c *Module_instance_identifierContext)

	// EnterNet_identifier is called when entering the net_identifier production.
	EnterNet_identifier(c *Net_identifierContext)

	// EnterOutput_port_identifier is called when entering the output_port_identifier production.
	EnterOutput_port_identifier(c *Output_port_identifierContext)

	// EnterParameter_identifier is called when entering the parameter_identifier production.
	EnterParameter_identifier(c *Parameter_identifierContext)

	// EnterPort_identifier is called when entering the port_identifier production.
	EnterPort_identifier(c *Port_identifierContext)

	// EnterReal_identifier is called when entering the real_identifier production.
	EnterReal_identifier(c *Real_identifierContext)

	// EnterSimple_arrayed_identifier is called when entering the simple_arrayed_identifier production.
	EnterSimple_arrayed_identifier(c *Simple_arrayed_identifierContext)

	// EnterSimple_hierarchical_identifier is called when entering the simple_hierarchical_identifier production.
	EnterSimple_hierarchical_identifier(c *Simple_hierarchical_identifierContext)

	// EnterSpecparam_identifier is called when entering the specparam_identifier production.
	EnterSpecparam_identifier(c *Specparam_identifierContext)

	// EnterSystem_function_identifier is called when entering the system_function_identifier production.
	EnterSystem_function_identifier(c *System_function_identifierContext)

	// EnterSystem_task_identifier is called when entering the system_task_identifier production.
	EnterSystem_task_identifier(c *System_task_identifierContext)

	// EnterTask_identifier is called when entering the task_identifier production.
	EnterTask_identifier(c *Task_identifierContext)

	// EnterTerminal_identifier is called when entering the terminal_identifier production.
	EnterTerminal_identifier(c *Terminal_identifierContext)

	// EnterText_macro_identifier is called when entering the text_macro_identifier production.
	EnterText_macro_identifier(c *Text_macro_identifierContext)

	// EnterTopmodule_identifier is called when entering the topmodule_identifier production.
	EnterTopmodule_identifier(c *Topmodule_identifierContext)

	// EnterUdp_identifier is called when entering the udp_identifier production.
	EnterUdp_identifier(c *Udp_identifierContext)

	// EnterUdp_instance_identifier is called when entering the udp_instance_identifier production.
	EnterUdp_instance_identifier(c *Udp_instance_identifierContext)

	// EnterVariable_identifier is called when entering the variable_identifier production.
	EnterVariable_identifier(c *Variable_identifierContext)

	// EnterSimple_hierarchical_branch is called when entering the simple_hierarchical_branch production.
	EnterSimple_hierarchical_branch(c *Simple_hierarchical_branchContext)

	// EnterEscaped_hierarchical_branch is called when entering the escaped_hierarchical_branch production.
	EnterEscaped_hierarchical_branch(c *Escaped_hierarchical_branchContext)

	// ExitConfig_declaration is called when exiting the config_declaration production.
	ExitConfig_declaration(c *Config_declarationContext)

	// ExitDesign_statement is called when exiting the design_statement production.
	ExitDesign_statement(c *Design_statementContext)

	// ExitConfig_rule_statement is called when exiting the config_rule_statement production.
	ExitConfig_rule_statement(c *Config_rule_statementContext)

	// ExitDefault_clause is called when exiting the default_clause production.
	ExitDefault_clause(c *Default_clauseContext)

	// ExitInst_clause is called when exiting the inst_clause production.
	ExitInst_clause(c *Inst_clauseContext)

	// ExitInst_name is called when exiting the inst_name production.
	ExitInst_name(c *Inst_nameContext)

	// ExitLiblist_clause is called when exiting the liblist_clause production.
	ExitLiblist_clause(c *Liblist_clauseContext)

	// ExitCell_clause is called when exiting the cell_clause production.
	ExitCell_clause(c *Cell_clauseContext)

	// ExitUse_clause is called when exiting the use_clause production.
	ExitUse_clause(c *Use_clauseContext)

	// ExitSource_text is called when exiting the source_text production.
	ExitSource_text(c *Source_textContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)

	// ExitModule_declaration is called when exiting the module_declaration production.
	ExitModule_declaration(c *Module_declarationContext)

	// ExitModule_keyword is called when exiting the module_keyword production.
	ExitModule_keyword(c *Module_keywordContext)

	// ExitModule_parameter_port_list is called when exiting the module_parameter_port_list production.
	ExitModule_parameter_port_list(c *Module_parameter_port_listContext)

	// ExitList_of_ports is called when exiting the list_of_ports production.
	ExitList_of_ports(c *List_of_portsContext)

	// ExitList_of_port_declarations is called when exiting the list_of_port_declarations production.
	ExitList_of_port_declarations(c *List_of_port_declarationsContext)

	// ExitPort is called when exiting the port production.
	ExitPort(c *PortContext)

	// ExitPort_expression is called when exiting the port_expression production.
	ExitPort_expression(c *Port_expressionContext)

	// ExitPort_reference is called when exiting the port_reference production.
	ExitPort_reference(c *Port_referenceContext)

	// ExitPort_declaration is called when exiting the port_declaration production.
	ExitPort_declaration(c *Port_declarationContext)

	// ExitModule_item is called when exiting the module_item production.
	ExitModule_item(c *Module_itemContext)

	// ExitModule_or_generate_item is called when exiting the module_or_generate_item production.
	ExitModule_or_generate_item(c *Module_or_generate_itemContext)

	// ExitNon_port_module_item is called when exiting the non_port_module_item production.
	ExitNon_port_module_item(c *Non_port_module_itemContext)

	// ExitModule_or_generate_item_declaration is called when exiting the module_or_generate_item_declaration production.
	ExitModule_or_generate_item_declaration(c *Module_or_generate_item_declarationContext)

	// ExitParameter_override is called when exiting the parameter_override production.
	ExitParameter_override(c *Parameter_overrideContext)

	// ExitLocal_parameter_declaration is called when exiting the local_parameter_declaration production.
	ExitLocal_parameter_declaration(c *Local_parameter_declarationContext)

	// ExitParameter_declaration is called when exiting the parameter_declaration production.
	ExitParameter_declaration(c *Parameter_declarationContext)

	// ExitParameter_declaration_ is called when exiting the parameter_declaration_ production.
	ExitParameter_declaration_(c *Parameter_declaration_Context)

	// ExitSpecparam_declaration is called when exiting the specparam_declaration production.
	ExitSpecparam_declaration(c *Specparam_declarationContext)

	// ExitInout_declaration is called when exiting the inout_declaration production.
	ExitInout_declaration(c *Inout_declarationContext)

	// ExitInput_declaration is called when exiting the input_declaration production.
	ExitInput_declaration(c *Input_declarationContext)

	// ExitOutput_declaration is called when exiting the output_declaration production.
	ExitOutput_declaration(c *Output_declarationContext)

	// ExitEvent_declaration is called when exiting the event_declaration production.
	ExitEvent_declaration(c *Event_declarationContext)

	// ExitGenvar_declaration is called when exiting the genvar_declaration production.
	ExitGenvar_declaration(c *Genvar_declarationContext)

	// ExitInteger_declaration is called when exiting the integer_declaration production.
	ExitInteger_declaration(c *Integer_declarationContext)

	// ExitTime_declaration is called when exiting the time_declaration production.
	ExitTime_declaration(c *Time_declarationContext)

	// ExitReal_declaration is called when exiting the real_declaration production.
	ExitReal_declaration(c *Real_declarationContext)

	// ExitRealtime_declaration is called when exiting the realtime_declaration production.
	ExitRealtime_declaration(c *Realtime_declarationContext)

	// ExitReg_declaration is called when exiting the reg_declaration production.
	ExitReg_declaration(c *Reg_declarationContext)

	// ExitNet_declaration is called when exiting the net_declaration production.
	ExitNet_declaration(c *Net_declarationContext)

	// ExitNet_type is called when exiting the net_type production.
	ExitNet_type(c *Net_typeContext)

	// ExitOutput_variable_type is called when exiting the output_variable_type production.
	ExitOutput_variable_type(c *Output_variable_typeContext)

	// ExitReal_type is called when exiting the real_type production.
	ExitReal_type(c *Real_typeContext)

	// ExitVariable_type is called when exiting the variable_type production.
	ExitVariable_type(c *Variable_typeContext)

	// ExitDrive_strength is called when exiting the drive_strength production.
	ExitDrive_strength(c *Drive_strengthContext)

	// ExitStrength0 is called when exiting the strength0 production.
	ExitStrength0(c *Strength0Context)

	// ExitStrength1 is called when exiting the strength1 production.
	ExitStrength1(c *Strength1Context)

	// ExitCharge_strength is called when exiting the charge_strength production.
	ExitCharge_strength(c *Charge_strengthContext)

	// ExitDelay3 is called when exiting the delay3 production.
	ExitDelay3(c *Delay3Context)

	// ExitDelay2 is called when exiting the delay2 production.
	ExitDelay2(c *Delay2Context)

	// ExitDelay_value is called when exiting the delay_value production.
	ExitDelay_value(c *Delay_valueContext)

	// ExitList_of_event_identifiers is called when exiting the list_of_event_identifiers production.
	ExitList_of_event_identifiers(c *List_of_event_identifiersContext)

	// ExitList_of_net_identifiers is called when exiting the list_of_net_identifiers production.
	ExitList_of_net_identifiers(c *List_of_net_identifiersContext)

	// ExitList_of_genvar_identifiers is called when exiting the list_of_genvar_identifiers production.
	ExitList_of_genvar_identifiers(c *List_of_genvar_identifiersContext)

	// ExitList_of_port_identifiers is called when exiting the list_of_port_identifiers production.
	ExitList_of_port_identifiers(c *List_of_port_identifiersContext)

	// ExitList_of_net_decl_assignments is called when exiting the list_of_net_decl_assignments production.
	ExitList_of_net_decl_assignments(c *List_of_net_decl_assignmentsContext)

	// ExitList_of_param_assignments is called when exiting the list_of_param_assignments production.
	ExitList_of_param_assignments(c *List_of_param_assignmentsContext)

	// ExitList_of_specparam_assignments is called when exiting the list_of_specparam_assignments production.
	ExitList_of_specparam_assignments(c *List_of_specparam_assignmentsContext)

	// ExitList_of_real_identifiers is called when exiting the list_of_real_identifiers production.
	ExitList_of_real_identifiers(c *List_of_real_identifiersContext)

	// ExitList_of_variable_identifiers is called when exiting the list_of_variable_identifiers production.
	ExitList_of_variable_identifiers(c *List_of_variable_identifiersContext)

	// ExitList_of_variable_port_identifiers is called when exiting the list_of_variable_port_identifiers production.
	ExitList_of_variable_port_identifiers(c *List_of_variable_port_identifiersContext)

	// ExitNet_decl_assignment is called when exiting the net_decl_assignment production.
	ExitNet_decl_assignment(c *Net_decl_assignmentContext)

	// ExitParam_assignment is called when exiting the param_assignment production.
	ExitParam_assignment(c *Param_assignmentContext)

	// ExitSpecparam_assignment is called when exiting the specparam_assignment production.
	ExitSpecparam_assignment(c *Specparam_assignmentContext)

	// ExitPulse_control_specparam is called when exiting the pulse_control_specparam production.
	ExitPulse_control_specparam(c *Pulse_control_specparamContext)

	// ExitError_limit_value is called when exiting the error_limit_value production.
	ExitError_limit_value(c *Error_limit_valueContext)

	// ExitReject_limit_value is called when exiting the reject_limit_value production.
	ExitReject_limit_value(c *Reject_limit_valueContext)

	// ExitLimit_value is called when exiting the limit_value production.
	ExitLimit_value(c *Limit_valueContext)

	// ExitDimension is called when exiting the dimension production.
	ExitDimension(c *DimensionContext)

	// ExitRange_ is called when exiting the range_ production.
	ExitRange_(c *Range_Context)

	// ExitFunction_declaration is called when exiting the function_declaration production.
	ExitFunction_declaration(c *Function_declarationContext)

	// ExitFunction_item_declaration is called when exiting the function_item_declaration production.
	ExitFunction_item_declaration(c *Function_item_declarationContext)

	// ExitFunction_port_list is called when exiting the function_port_list production.
	ExitFunction_port_list(c *Function_port_listContext)

	// ExitFunction_port is called when exiting the function_port production.
	ExitFunction_port(c *Function_portContext)

	// ExitRange_or_type is called when exiting the range_or_type production.
	ExitRange_or_type(c *Range_or_typeContext)

	// ExitTask_declaration is called when exiting the task_declaration production.
	ExitTask_declaration(c *Task_declarationContext)

	// ExitTask_item_declaration is called when exiting the task_item_declaration production.
	ExitTask_item_declaration(c *Task_item_declarationContext)

	// ExitTask_port_list is called when exiting the task_port_list production.
	ExitTask_port_list(c *Task_port_listContext)

	// ExitTask_port_item is called when exiting the task_port_item production.
	ExitTask_port_item(c *Task_port_itemContext)

	// ExitTf_decl_header is called when exiting the tf_decl_header production.
	ExitTf_decl_header(c *Tf_decl_headerContext)

	// ExitTf_declaration is called when exiting the tf_declaration production.
	ExitTf_declaration(c *Tf_declarationContext)

	// ExitTask_port_type is called when exiting the task_port_type production.
	ExitTask_port_type(c *Task_port_typeContext)

	// ExitBlock_item_declaration is called when exiting the block_item_declaration production.
	ExitBlock_item_declaration(c *Block_item_declarationContext)

	// ExitBlock_reg_declaration is called when exiting the block_reg_declaration production.
	ExitBlock_reg_declaration(c *Block_reg_declarationContext)

	// ExitList_of_block_variable_identifiers is called when exiting the list_of_block_variable_identifiers production.
	ExitList_of_block_variable_identifiers(c *List_of_block_variable_identifiersContext)

	// ExitBlock_variable_type is called when exiting the block_variable_type production.
	ExitBlock_variable_type(c *Block_variable_typeContext)

	// ExitGate_instantiation is called when exiting the gate_instantiation production.
	ExitGate_instantiation(c *Gate_instantiationContext)

	// ExitCmos_switch_instance is called when exiting the cmos_switch_instance production.
	ExitCmos_switch_instance(c *Cmos_switch_instanceContext)

	// ExitEnable_gate_instance is called when exiting the enable_gate_instance production.
	ExitEnable_gate_instance(c *Enable_gate_instanceContext)

	// ExitMos_switch_instance is called when exiting the mos_switch_instance production.
	ExitMos_switch_instance(c *Mos_switch_instanceContext)

	// ExitN_input_gate_instance is called when exiting the n_input_gate_instance production.
	ExitN_input_gate_instance(c *N_input_gate_instanceContext)

	// ExitN_output_gate_instance is called when exiting the n_output_gate_instance production.
	ExitN_output_gate_instance(c *N_output_gate_instanceContext)

	// ExitPass_switch_instance is called when exiting the pass_switch_instance production.
	ExitPass_switch_instance(c *Pass_switch_instanceContext)

	// ExitPass_enable_switch_instance is called when exiting the pass_enable_switch_instance production.
	ExitPass_enable_switch_instance(c *Pass_enable_switch_instanceContext)

	// ExitPull_gate_instance is called when exiting the pull_gate_instance production.
	ExitPull_gate_instance(c *Pull_gate_instanceContext)

	// ExitName_of_gate_instance is called when exiting the name_of_gate_instance production.
	ExitName_of_gate_instance(c *Name_of_gate_instanceContext)

	// ExitPulldown_strength is called when exiting the pulldown_strength production.
	ExitPulldown_strength(c *Pulldown_strengthContext)

	// ExitPullup_strength is called when exiting the pullup_strength production.
	ExitPullup_strength(c *Pullup_strengthContext)

	// ExitEnable_terminal is called when exiting the enable_terminal production.
	ExitEnable_terminal(c *Enable_terminalContext)

	// ExitNcontrol_terminal is called when exiting the ncontrol_terminal production.
	ExitNcontrol_terminal(c *Ncontrol_terminalContext)

	// ExitPcontrol_terminal is called when exiting the pcontrol_terminal production.
	ExitPcontrol_terminal(c *Pcontrol_terminalContext)

	// ExitInput_terminal is called when exiting the input_terminal production.
	ExitInput_terminal(c *Input_terminalContext)

	// ExitInout_terminal is called when exiting the inout_terminal production.
	ExitInout_terminal(c *Inout_terminalContext)

	// ExitOutput_terminal is called when exiting the output_terminal production.
	ExitOutput_terminal(c *Output_terminalContext)

	// ExitCmos_switchtype is called when exiting the cmos_switchtype production.
	ExitCmos_switchtype(c *Cmos_switchtypeContext)

	// ExitEnable_gatetype is called when exiting the enable_gatetype production.
	ExitEnable_gatetype(c *Enable_gatetypeContext)

	// ExitMos_switchtype is called when exiting the mos_switchtype production.
	ExitMos_switchtype(c *Mos_switchtypeContext)

	// ExitN_input_gatetype is called when exiting the n_input_gatetype production.
	ExitN_input_gatetype(c *N_input_gatetypeContext)

	// ExitN_output_gatetype is called when exiting the n_output_gatetype production.
	ExitN_output_gatetype(c *N_output_gatetypeContext)

	// ExitPass_en_switchtype is called when exiting the pass_en_switchtype production.
	ExitPass_en_switchtype(c *Pass_en_switchtypeContext)

	// ExitPass_switchtype is called when exiting the pass_switchtype production.
	ExitPass_switchtype(c *Pass_switchtypeContext)

	// ExitModule_instantiation is called when exiting the module_instantiation production.
	ExitModule_instantiation(c *Module_instantiationContext)

	// ExitParameter_value_assignment is called when exiting the parameter_value_assignment production.
	ExitParameter_value_assignment(c *Parameter_value_assignmentContext)

	// ExitList_of_parameter_assignments is called when exiting the list_of_parameter_assignments production.
	ExitList_of_parameter_assignments(c *List_of_parameter_assignmentsContext)

	// ExitOrdered_parameter_assignment is called when exiting the ordered_parameter_assignment production.
	ExitOrdered_parameter_assignment(c *Ordered_parameter_assignmentContext)

	// ExitNamed_parameter_assignment is called when exiting the named_parameter_assignment production.
	ExitNamed_parameter_assignment(c *Named_parameter_assignmentContext)

	// ExitModule_instance is called when exiting the module_instance production.
	ExitModule_instance(c *Module_instanceContext)

	// ExitName_of_instance is called when exiting the name_of_instance production.
	ExitName_of_instance(c *Name_of_instanceContext)

	// ExitList_of_port_connections is called when exiting the list_of_port_connections production.
	ExitList_of_port_connections(c *List_of_port_connectionsContext)

	// ExitOrdered_port_connection is called when exiting the ordered_port_connection production.
	ExitOrdered_port_connection(c *Ordered_port_connectionContext)

	// ExitNamed_port_connection is called when exiting the named_port_connection production.
	ExitNamed_port_connection(c *Named_port_connectionContext)

	// ExitGenerated_instantiation is called when exiting the generated_instantiation production.
	ExitGenerated_instantiation(c *Generated_instantiationContext)

	// ExitGenerate_item_or_null is called when exiting the generate_item_or_null production.
	ExitGenerate_item_or_null(c *Generate_item_or_nullContext)

	// ExitGenerate_item is called when exiting the generate_item production.
	ExitGenerate_item(c *Generate_itemContext)

	// ExitGenerate_conditional_statement is called when exiting the generate_conditional_statement production.
	ExitGenerate_conditional_statement(c *Generate_conditional_statementContext)

	// ExitGenerate_case_statement is called when exiting the generate_case_statement production.
	ExitGenerate_case_statement(c *Generate_case_statementContext)

	// ExitGenvar_case_item is called when exiting the genvar_case_item production.
	ExitGenvar_case_item(c *Genvar_case_itemContext)

	// ExitGenerate_loop_statement is called when exiting the generate_loop_statement production.
	ExitGenerate_loop_statement(c *Generate_loop_statementContext)

	// ExitGenvar_assignment is called when exiting the genvar_assignment production.
	ExitGenvar_assignment(c *Genvar_assignmentContext)

	// ExitGenerate_block is called when exiting the generate_block production.
	ExitGenerate_block(c *Generate_blockContext)

	// ExitContinuous_assign is called when exiting the continuous_assign production.
	ExitContinuous_assign(c *Continuous_assignContext)

	// ExitList_of_net_assignments is called when exiting the list_of_net_assignments production.
	ExitList_of_net_assignments(c *List_of_net_assignmentsContext)

	// ExitNet_assignment is called when exiting the net_assignment production.
	ExitNet_assignment(c *Net_assignmentContext)

	// ExitInitial_construct is called when exiting the initial_construct production.
	ExitInitial_construct(c *Initial_constructContext)

	// ExitAlways_construct is called when exiting the always_construct production.
	ExitAlways_construct(c *Always_constructContext)

	// ExitBlocking_assignment is called when exiting the blocking_assignment production.
	ExitBlocking_assignment(c *Blocking_assignmentContext)

	// ExitNonblocking_assignment is called when exiting the nonblocking_assignment production.
	ExitNonblocking_assignment(c *Nonblocking_assignmentContext)

	// ExitProcedural_continuous_assignments is called when exiting the procedural_continuous_assignments production.
	ExitProcedural_continuous_assignments(c *Procedural_continuous_assignmentsContext)

	// ExitFunction_blocking_assignment is called when exiting the function_blocking_assignment production.
	ExitFunction_blocking_assignment(c *Function_blocking_assignmentContext)

	// ExitFunction_statement_or_null is called when exiting the function_statement_or_null production.
	ExitFunction_statement_or_null(c *Function_statement_or_nullContext)

	// ExitFunction_seq_block is called when exiting the function_seq_block production.
	ExitFunction_seq_block(c *Function_seq_blockContext)

	// ExitVariable_assignment is called when exiting the variable_assignment production.
	ExitVariable_assignment(c *Variable_assignmentContext)

	// ExitPar_block is called when exiting the par_block production.
	ExitPar_block(c *Par_blockContext)

	// ExitSeq_block is called when exiting the seq_block production.
	ExitSeq_block(c *Seq_blockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitStatement_or_null is called when exiting the statement_or_null production.
	ExitStatement_or_null(c *Statement_or_nullContext)

	// ExitFunction_statement is called when exiting the function_statement production.
	ExitFunction_statement(c *Function_statementContext)

	// ExitDelay_or_event_control is called when exiting the delay_or_event_control production.
	ExitDelay_or_event_control(c *Delay_or_event_controlContext)

	// ExitDelay_control is called when exiting the delay_control production.
	ExitDelay_control(c *Delay_controlContext)

	// ExitDisable_statement is called when exiting the disable_statement production.
	ExitDisable_statement(c *Disable_statementContext)

	// ExitEvent_control is called when exiting the event_control production.
	ExitEvent_control(c *Event_controlContext)

	// ExitEvent_trigger is called when exiting the event_trigger production.
	ExitEvent_trigger(c *Event_triggerContext)

	// ExitEvent_expression is called when exiting the event_expression production.
	ExitEvent_expression(c *Event_expressionContext)

	// ExitEvent_primary is called when exiting the event_primary production.
	ExitEvent_primary(c *Event_primaryContext)

	// ExitProcedural_timing_control_statement is called when exiting the procedural_timing_control_statement production.
	ExitProcedural_timing_control_statement(c *Procedural_timing_control_statementContext)

	// ExitWait_statement is called when exiting the wait_statement production.
	ExitWait_statement(c *Wait_statementContext)

	// ExitConditional_statement is called when exiting the conditional_statement production.
	ExitConditional_statement(c *Conditional_statementContext)

	// ExitIf_else_if_statement is called when exiting the if_else_if_statement production.
	ExitIf_else_if_statement(c *If_else_if_statementContext)

	// ExitElse_statement is called when exiting the else_statement production.
	ExitElse_statement(c *Else_statementContext)

	// ExitFunction_conditional_statement is called when exiting the function_conditional_statement production.
	ExitFunction_conditional_statement(c *Function_conditional_statementContext)

	// ExitFunction_if_else_if_statement is called when exiting the function_if_else_if_statement production.
	ExitFunction_if_else_if_statement(c *Function_if_else_if_statementContext)

	// ExitCase_statement is called when exiting the case_statement production.
	ExitCase_statement(c *Case_statementContext)

	// ExitCase_item is called when exiting the case_item production.
	ExitCase_item(c *Case_itemContext)

	// ExitFunction_case_statement is called when exiting the function_case_statement production.
	ExitFunction_case_statement(c *Function_case_statementContext)

	// ExitFunction_case_item is called when exiting the function_case_item production.
	ExitFunction_case_item(c *Function_case_itemContext)

	// ExitFunction_loop_statement is called when exiting the function_loop_statement production.
	ExitFunction_loop_statement(c *Function_loop_statementContext)

	// ExitLoop_statement is called when exiting the loop_statement production.
	ExitLoop_statement(c *Loop_statementContext)

	// ExitSystem_task_enable is called when exiting the system_task_enable production.
	ExitSystem_task_enable(c *System_task_enableContext)

	// ExitTask_enable is called when exiting the task_enable production.
	ExitTask_enable(c *Task_enableContext)

	// ExitSpecify_block is called when exiting the specify_block production.
	ExitSpecify_block(c *Specify_blockContext)

	// ExitSpecify_item is called when exiting the specify_item production.
	ExitSpecify_item(c *Specify_itemContext)

	// ExitPulsestyle_declaration is called when exiting the pulsestyle_declaration production.
	ExitPulsestyle_declaration(c *Pulsestyle_declarationContext)

	// ExitShowcancelled_declaration is called when exiting the showcancelled_declaration production.
	ExitShowcancelled_declaration(c *Showcancelled_declarationContext)

	// ExitPath_declaration is called when exiting the path_declaration production.
	ExitPath_declaration(c *Path_declarationContext)

	// ExitSimple_path_declaration is called when exiting the simple_path_declaration production.
	ExitSimple_path_declaration(c *Simple_path_declarationContext)

	// ExitParallel_path_description is called when exiting the parallel_path_description production.
	ExitParallel_path_description(c *Parallel_path_descriptionContext)

	// ExitFull_path_description is called when exiting the full_path_description production.
	ExitFull_path_description(c *Full_path_descriptionContext)

	// ExitList_of_path_inputs is called when exiting the list_of_path_inputs production.
	ExitList_of_path_inputs(c *List_of_path_inputsContext)

	// ExitList_of_path_outputs is called when exiting the list_of_path_outputs production.
	ExitList_of_path_outputs(c *List_of_path_outputsContext)

	// ExitSpecify_input_terminal_descriptor is called when exiting the specify_input_terminal_descriptor production.
	ExitSpecify_input_terminal_descriptor(c *Specify_input_terminal_descriptorContext)

	// ExitSpecify_output_terminal_descriptor is called when exiting the specify_output_terminal_descriptor production.
	ExitSpecify_output_terminal_descriptor(c *Specify_output_terminal_descriptorContext)

	// ExitInput_identifier is called when exiting the input_identifier production.
	ExitInput_identifier(c *Input_identifierContext)

	// ExitOutput_identifier is called when exiting the output_identifier production.
	ExitOutput_identifier(c *Output_identifierContext)

	// ExitPath_delay_value is called when exiting the path_delay_value production.
	ExitPath_delay_value(c *Path_delay_valueContext)

	// ExitList_of_path_delay_expressions is called when exiting the list_of_path_delay_expressions production.
	ExitList_of_path_delay_expressions(c *List_of_path_delay_expressionsContext)

	// ExitT_path_delay_expression is called when exiting the t_path_delay_expression production.
	ExitT_path_delay_expression(c *T_path_delay_expressionContext)

	// ExitTrise_path_delay_expression is called when exiting the trise_path_delay_expression production.
	ExitTrise_path_delay_expression(c *Trise_path_delay_expressionContext)

	// ExitTfall_path_delay_expression is called when exiting the tfall_path_delay_expression production.
	ExitTfall_path_delay_expression(c *Tfall_path_delay_expressionContext)

	// ExitTz_path_delay_expression is called when exiting the tz_path_delay_expression production.
	ExitTz_path_delay_expression(c *Tz_path_delay_expressionContext)

	// ExitT01_path_delay_expression is called when exiting the t01_path_delay_expression production.
	ExitT01_path_delay_expression(c *T01_path_delay_expressionContext)

	// ExitT10_path_delay_expression is called when exiting the t10_path_delay_expression production.
	ExitT10_path_delay_expression(c *T10_path_delay_expressionContext)

	// ExitT0z_path_delay_expression is called when exiting the t0z_path_delay_expression production.
	ExitT0z_path_delay_expression(c *T0z_path_delay_expressionContext)

	// ExitTz1_path_delay_expression is called when exiting the tz1_path_delay_expression production.
	ExitTz1_path_delay_expression(c *Tz1_path_delay_expressionContext)

	// ExitT1z_path_delay_expression is called when exiting the t1z_path_delay_expression production.
	ExitT1z_path_delay_expression(c *T1z_path_delay_expressionContext)

	// ExitTz0_path_delay_expression is called when exiting the tz0_path_delay_expression production.
	ExitTz0_path_delay_expression(c *Tz0_path_delay_expressionContext)

	// ExitT0x_path_delay_expression is called when exiting the t0x_path_delay_expression production.
	ExitT0x_path_delay_expression(c *T0x_path_delay_expressionContext)

	// ExitTx1_path_delay_expression is called when exiting the tx1_path_delay_expression production.
	ExitTx1_path_delay_expression(c *Tx1_path_delay_expressionContext)

	// ExitT1x_path_delay_expression is called when exiting the t1x_path_delay_expression production.
	ExitT1x_path_delay_expression(c *T1x_path_delay_expressionContext)

	// ExitTx0_path_delay_expression is called when exiting the tx0_path_delay_expression production.
	ExitTx0_path_delay_expression(c *Tx0_path_delay_expressionContext)

	// ExitTxz_path_delay_expression is called when exiting the txz_path_delay_expression production.
	ExitTxz_path_delay_expression(c *Txz_path_delay_expressionContext)

	// ExitTzx_path_delay_expression is called when exiting the tzx_path_delay_expression production.
	ExitTzx_path_delay_expression(c *Tzx_path_delay_expressionContext)

	// ExitPath_delay_expression is called when exiting the path_delay_expression production.
	ExitPath_delay_expression(c *Path_delay_expressionContext)

	// ExitEdge_sensitive_path_declaration is called when exiting the edge_sensitive_path_declaration production.
	ExitEdge_sensitive_path_declaration(c *Edge_sensitive_path_declarationContext)

	// ExitParallel_edge_sensitive_path_description is called when exiting the parallel_edge_sensitive_path_description production.
	ExitParallel_edge_sensitive_path_description(c *Parallel_edge_sensitive_path_descriptionContext)

	// ExitFull_edge_sensitive_path_description is called when exiting the full_edge_sensitive_path_description production.
	ExitFull_edge_sensitive_path_description(c *Full_edge_sensitive_path_descriptionContext)

	// ExitData_source_expression is called when exiting the data_source_expression production.
	ExitData_source_expression(c *Data_source_expressionContext)

	// ExitEdge_identifier is called when exiting the edge_identifier production.
	ExitEdge_identifier(c *Edge_identifierContext)

	// ExitState_dependent_path_declaration is called when exiting the state_dependent_path_declaration production.
	ExitState_dependent_path_declaration(c *State_dependent_path_declarationContext)

	// ExitPolarity_operator is called when exiting the polarity_operator production.
	ExitPolarity_operator(c *Polarity_operatorContext)

	// ExitChecktime_condition is called when exiting the checktime_condition production.
	ExitChecktime_condition(c *Checktime_conditionContext)

	// ExitDelayed_data is called when exiting the delayed_data production.
	ExitDelayed_data(c *Delayed_dataContext)

	// ExitDelayed_reference is called when exiting the delayed_reference production.
	ExitDelayed_reference(c *Delayed_referenceContext)

	// ExitEnd_edge_offset is called when exiting the end_edge_offset production.
	ExitEnd_edge_offset(c *End_edge_offsetContext)

	// ExitEvent_based_flag is called when exiting the event_based_flag production.
	ExitEvent_based_flag(c *Event_based_flagContext)

	// ExitNotify_reg is called when exiting the notify_reg production.
	ExitNotify_reg(c *Notify_regContext)

	// ExitRemain_active_flag is called when exiting the remain_active_flag production.
	ExitRemain_active_flag(c *Remain_active_flagContext)

	// ExitStamptime_condition is called when exiting the stamptime_condition production.
	ExitStamptime_condition(c *Stamptime_conditionContext)

	// ExitStart_edge_offset is called when exiting the start_edge_offset production.
	ExitStart_edge_offset(c *Start_edge_offsetContext)

	// ExitThreshold is called when exiting the threshold production.
	ExitThreshold(c *ThresholdContext)

	// ExitTiming_check_limit is called when exiting the timing_check_limit production.
	ExitTiming_check_limit(c *Timing_check_limitContext)

	// ExitConcatenation is called when exiting the concatenation production.
	ExitConcatenation(c *ConcatenationContext)

	// ExitConstant_concatenation is called when exiting the constant_concatenation production.
	ExitConstant_concatenation(c *Constant_concatenationContext)

	// ExitConstant_multiple_concatenation is called when exiting the constant_multiple_concatenation production.
	ExitConstant_multiple_concatenation(c *Constant_multiple_concatenationContext)

	// ExitModule_path_concatenation is called when exiting the module_path_concatenation production.
	ExitModule_path_concatenation(c *Module_path_concatenationContext)

	// ExitModule_path_multiple_concatenation is called when exiting the module_path_multiple_concatenation production.
	ExitModule_path_multiple_concatenation(c *Module_path_multiple_concatenationContext)

	// ExitMultiple_concatenation is called when exiting the multiple_concatenation production.
	ExitMultiple_concatenation(c *Multiple_concatenationContext)

	// ExitNet_concatenation is called when exiting the net_concatenation production.
	ExitNet_concatenation(c *Net_concatenationContext)

	// ExitNet_concatenation_value is called when exiting the net_concatenation_value production.
	ExitNet_concatenation_value(c *Net_concatenation_valueContext)

	// ExitVariable_concatenation is called when exiting the variable_concatenation production.
	ExitVariable_concatenation(c *Variable_concatenationContext)

	// ExitVariable_concatenation_value is called when exiting the variable_concatenation_value production.
	ExitVariable_concatenation_value(c *Variable_concatenation_valueContext)

	// ExitConstant_function_call is called when exiting the constant_function_call production.
	ExitConstant_function_call(c *Constant_function_callContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitSystem_function_call is called when exiting the system_function_call production.
	ExitSystem_function_call(c *System_function_callContext)

	// ExitGenvar_function_call is called when exiting the genvar_function_call production.
	ExitGenvar_function_call(c *Genvar_function_callContext)

	// ExitBase_expression is called when exiting the base_expression production.
	ExitBase_expression(c *Base_expressionContext)

	// ExitConstant_base_expression is called when exiting the constant_base_expression production.
	ExitConstant_base_expression(c *Constant_base_expressionContext)

	// ExitConstant_expression is called when exiting the constant_expression production.
	ExitConstant_expression(c *Constant_expressionContext)

	// ExitConstant_mintypmax_expression is called when exiting the constant_mintypmax_expression production.
	ExitConstant_mintypmax_expression(c *Constant_mintypmax_expressionContext)

	// ExitConstant_range_expression is called when exiting the constant_range_expression production.
	ExitConstant_range_expression(c *Constant_range_expressionContext)

	// ExitDimension_constant_expression is called when exiting the dimension_constant_expression production.
	ExitDimension_constant_expression(c *Dimension_constant_expressionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitLsb_constant_expression is called when exiting the lsb_constant_expression production.
	ExitLsb_constant_expression(c *Lsb_constant_expressionContext)

	// ExitMintypmax_expression is called when exiting the mintypmax_expression production.
	ExitMintypmax_expression(c *Mintypmax_expressionContext)

	// ExitModule_path_conditional_expression is called when exiting the module_path_conditional_expression production.
	ExitModule_path_conditional_expression(c *Module_path_conditional_expressionContext)

	// ExitModule_path_expression is called when exiting the module_path_expression production.
	ExitModule_path_expression(c *Module_path_expressionContext)

	// ExitModule_path_mintypmax_expression is called when exiting the module_path_mintypmax_expression production.
	ExitModule_path_mintypmax_expression(c *Module_path_mintypmax_expressionContext)

	// ExitMsb_constant_expression is called when exiting the msb_constant_expression production.
	ExitMsb_constant_expression(c *Msb_constant_expressionContext)

	// ExitRange_expression is called when exiting the range_expression production.
	ExitRange_expression(c *Range_expressionContext)

	// ExitWidth_constant_expression is called when exiting the width_constant_expression production.
	ExitWidth_constant_expression(c *Width_constant_expressionContext)

	// ExitConstant_primary is called when exiting the constant_primary production.
	ExitConstant_primary(c *Constant_primaryContext)

	// ExitModule_path_primary is called when exiting the module_path_primary production.
	ExitModule_path_primary(c *Module_path_primaryContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitNet_lvalue is called when exiting the net_lvalue production.
	ExitNet_lvalue(c *Net_lvalueContext)

	// ExitVariable_lvalue is called when exiting the variable_lvalue production.
	ExitVariable_lvalue(c *Variable_lvalueContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitBinary_operator is called when exiting the binary_operator production.
	ExitBinary_operator(c *Binary_operatorContext)

	// ExitUnary_module_path_operator is called when exiting the unary_module_path_operator production.
	ExitUnary_module_path_operator(c *Unary_module_path_operatorContext)

	// ExitBinary_module_path_operator is called when exiting the binary_module_path_operator production.
	ExitBinary_module_path_operator(c *Binary_module_path_operatorContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitTiming_spec is called when exiting the timing_spec production.
	ExitTiming_spec(c *Timing_specContext)

	// ExitDefine is called when exiting the define production.
	ExitDefine(c *DefineContext)

	// ExitAttribute_instance is called when exiting the attribute_instance production.
	ExitAttribute_instance(c *Attribute_instanceContext)

	// ExitAttr_spec is called when exiting the attr_spec production.
	ExitAttr_spec(c *Attr_specContext)

	// ExitAttr_name is called when exiting the attr_name production.
	ExitAttr_name(c *Attr_nameContext)

	// ExitArrayed_identifier is called when exiting the arrayed_identifier production.
	ExitArrayed_identifier(c *Arrayed_identifierContext)

	// ExitBlock_identifier is called when exiting the block_identifier production.
	ExitBlock_identifier(c *Block_identifierContext)

	// ExitCell_identifier is called when exiting the cell_identifier production.
	ExitCell_identifier(c *Cell_identifierContext)

	// ExitConfig_identifier is called when exiting the config_identifier production.
	ExitConfig_identifier(c *Config_identifierContext)

	// ExitDefine_identifier is called when exiting the define_identifier production.
	ExitDefine_identifier(c *Define_identifierContext)

	// ExitEscaped_arrayed_identifier is called when exiting the escaped_arrayed_identifier production.
	ExitEscaped_arrayed_identifier(c *Escaped_arrayed_identifierContext)

	// ExitEscaped_hierarchical_identifier is called when exiting the escaped_hierarchical_identifier production.
	ExitEscaped_hierarchical_identifier(c *Escaped_hierarchical_identifierContext)

	// ExitEvent_identifier is called when exiting the event_identifier production.
	ExitEvent_identifier(c *Event_identifierContext)

	// ExitFunction_identifier is called when exiting the function_identifier production.
	ExitFunction_identifier(c *Function_identifierContext)

	// ExitGate_instance_identifier is called when exiting the gate_instance_identifier production.
	ExitGate_instance_identifier(c *Gate_instance_identifierContext)

	// ExitGenerate_block_identifier is called when exiting the generate_block_identifier production.
	ExitGenerate_block_identifier(c *Generate_block_identifierContext)

	// ExitGenvar_function_identifier is called when exiting the genvar_function_identifier production.
	ExitGenvar_function_identifier(c *Genvar_function_identifierContext)

	// ExitGenvar_identifier is called when exiting the genvar_identifier production.
	ExitGenvar_identifier(c *Genvar_identifierContext)

	// ExitHierarchical_block_identifier is called when exiting the hierarchical_block_identifier production.
	ExitHierarchical_block_identifier(c *Hierarchical_block_identifierContext)

	// ExitHierarchical_event_identifier is called when exiting the hierarchical_event_identifier production.
	ExitHierarchical_event_identifier(c *Hierarchical_event_identifierContext)

	// ExitHierarchical_function_identifier is called when exiting the hierarchical_function_identifier production.
	ExitHierarchical_function_identifier(c *Hierarchical_function_identifierContext)

	// ExitHierarchical_identifier is called when exiting the hierarchical_identifier production.
	ExitHierarchical_identifier(c *Hierarchical_identifierContext)

	// ExitHierarchical_net_identifier is called when exiting the hierarchical_net_identifier production.
	ExitHierarchical_net_identifier(c *Hierarchical_net_identifierContext)

	// ExitHierarchical_variable_identifier is called when exiting the hierarchical_variable_identifier production.
	ExitHierarchical_variable_identifier(c *Hierarchical_variable_identifierContext)

	// ExitHierarchical_task_identifier is called when exiting the hierarchical_task_identifier production.
	ExitHierarchical_task_identifier(c *Hierarchical_task_identifierContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitInout_port_identifier is called when exiting the inout_port_identifier production.
	ExitInout_port_identifier(c *Inout_port_identifierContext)

	// ExitInput_port_identifier is called when exiting the input_port_identifier production.
	ExitInput_port_identifier(c *Input_port_identifierContext)

	// ExitInstance_identifier is called when exiting the instance_identifier production.
	ExitInstance_identifier(c *Instance_identifierContext)

	// ExitLibrary_identifier is called when exiting the library_identifier production.
	ExitLibrary_identifier(c *Library_identifierContext)

	// ExitMemory_identifier is called when exiting the memory_identifier production.
	ExitMemory_identifier(c *Memory_identifierContext)

	// ExitModule_identifier is called when exiting the module_identifier production.
	ExitModule_identifier(c *Module_identifierContext)

	// ExitModule_instance_identifier is called when exiting the module_instance_identifier production.
	ExitModule_instance_identifier(c *Module_instance_identifierContext)

	// ExitNet_identifier is called when exiting the net_identifier production.
	ExitNet_identifier(c *Net_identifierContext)

	// ExitOutput_port_identifier is called when exiting the output_port_identifier production.
	ExitOutput_port_identifier(c *Output_port_identifierContext)

	// ExitParameter_identifier is called when exiting the parameter_identifier production.
	ExitParameter_identifier(c *Parameter_identifierContext)

	// ExitPort_identifier is called when exiting the port_identifier production.
	ExitPort_identifier(c *Port_identifierContext)

	// ExitReal_identifier is called when exiting the real_identifier production.
	ExitReal_identifier(c *Real_identifierContext)

	// ExitSimple_arrayed_identifier is called when exiting the simple_arrayed_identifier production.
	ExitSimple_arrayed_identifier(c *Simple_arrayed_identifierContext)

	// ExitSimple_hierarchical_identifier is called when exiting the simple_hierarchical_identifier production.
	ExitSimple_hierarchical_identifier(c *Simple_hierarchical_identifierContext)

	// ExitSpecparam_identifier is called when exiting the specparam_identifier production.
	ExitSpecparam_identifier(c *Specparam_identifierContext)

	// ExitSystem_function_identifier is called when exiting the system_function_identifier production.
	ExitSystem_function_identifier(c *System_function_identifierContext)

	// ExitSystem_task_identifier is called when exiting the system_task_identifier production.
	ExitSystem_task_identifier(c *System_task_identifierContext)

	// ExitTask_identifier is called when exiting the task_identifier production.
	ExitTask_identifier(c *Task_identifierContext)

	// ExitTerminal_identifier is called when exiting the terminal_identifier production.
	ExitTerminal_identifier(c *Terminal_identifierContext)

	// ExitText_macro_identifier is called when exiting the text_macro_identifier production.
	ExitText_macro_identifier(c *Text_macro_identifierContext)

	// ExitTopmodule_identifier is called when exiting the topmodule_identifier production.
	ExitTopmodule_identifier(c *Topmodule_identifierContext)

	// ExitUdp_identifier is called when exiting the udp_identifier production.
	ExitUdp_identifier(c *Udp_identifierContext)

	// ExitUdp_instance_identifier is called when exiting the udp_instance_identifier production.
	ExitUdp_instance_identifier(c *Udp_instance_identifierContext)

	// ExitVariable_identifier is called when exiting the variable_identifier production.
	ExitVariable_identifier(c *Variable_identifierContext)

	// ExitSimple_hierarchical_branch is called when exiting the simple_hierarchical_branch production.
	ExitSimple_hierarchical_branch(c *Simple_hierarchical_branchContext)

	// ExitEscaped_hierarchical_branch is called when exiting the escaped_hierarchical_branch production.
	ExitEscaped_hierarchical_branch(c *Escaped_hierarchical_branchContext)
}
