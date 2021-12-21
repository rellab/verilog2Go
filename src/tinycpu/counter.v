module counter(clk,reset,load,inc,d,q);
   
  input clk,reset,load,inc;
  input [15:0] d;
  output [15:0] q;
  reg [15:0] 	 q;

  always @(posedge clk or negedge reset)
    if(!reset) q <= 16'h0000;
    else if(load) q <= d;
    else if(inc) q <= q;

endmodule