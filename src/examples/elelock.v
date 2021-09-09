module elelock(clk, reset, tenkey, close, lock);
input clk, reset, close;
input [9: 0] tenkey;
output reg lock;

reg ke1, ke2;
reg [3:0] key [0: 3];
wire match, key_enbl;

//暗証番号設定
parameter SECRET_3 = 4'h5, SECRET_2 = 4'h9, SECRET_1 = 4'h6, SECRET_0 = 4'h3;

//暗証番号入力レジスタ
always@(posedge clk or posedge reset) begin
	if(reset == 1'b1)begin
		key[3] <= 4'b1111;
		key[2] <= 4'b1111;
		key[1] <= 4'b1111;
		key[0] <= 4'b1111;
	end
	else if(close == 1'b1)begin
		key[3] <= 4'b1111;
		key[2] <= 4'b1111;
		key[1] <= 4'b1111;
		key[0] <= 4'b1111;
	end
	else if(key_enbl == 1'b1)begin
		key[3] <= key[2];
		key[2] <= key[1];
		key[1] <= key[0];
		key[0] <= keyenc(tenkey);
	end
end

//テンキー入力チャッタ取り用
always@(posedge clk or posedge reset) begin
	if(reset == 1'b1) begin
		ke2 <= 1'b0;
		ke1 <= 1'b0;
	end
	else begin
		ke2 <= ke1;
		ke1 <= | tenkey;
	end
end

//電子錠出力
always@(posedge clk or posedge reset) begin
	if(reset == 1'b1)
		lock <= 1'b1;
	else if(close == 1'b1)
		lock <= 1'b1;
	else if(match == 1'b1)
		lock <= 1'b0;
end

//テンキー入力エンコーダ
function [3:0] keyenc;
input [9:0] sw;
	case(sw)
		10'b0000000001: keyenc = 4'h0;
		10'b0000000010: keyenc = 4'h1;
		10'b0000000100: keyenc = 4'h2;
		10'b0000001000: keyenc = 4'h3;
		10'b0000010000: keyenc = 4'h4;
		10'b0000100000: keyenc = 4'h5;
		10'b0001000000: keyenc = 4'h6;
		10'b0010000000: keyenc = 4'h7;
		10'b0100000000: keyenc = 4'h8;
		10'b1000000000: keyenc = 4'h9;
	endcase
endfunction

//暗証番号一致信号
assign match = (key[0] == SECRET_0) && (key[1] == SECRET_1) && (key[2] == SECRET_2) && (key[3] == SECRET_3);
assign key_enbl = ~ke2 & ke1;

endmodule