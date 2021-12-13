module ram(clk, load, addr, d, q);
//  parameter DWIDTH=16,AWIDTH=12,WORDS=4096;

 input clk,load;
 input [11:0] addr;
 input [15:0] d;
 output [15:0] q;
 reg [15:0] q;
 reg [15:0] mem [4095:0];

 always @(posedge clk)
   begin
     if(load) mem[addr] <= d;
     q <= mem[addr];
   end

 integer i;
 initial begin
    for(i=0;i<WORDS;i=i+1)
       mem[i]=0;
 // ここにメモリの初期化（mem[12'h001]=16'h1234;など）を書く．
 end

endmodule