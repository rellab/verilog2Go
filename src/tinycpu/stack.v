module stack(clk, reset, load, push, pop, d, qtop, qnext);
  // parameter N = 8;
	 
  input clk, reset, load, push, pop;
  input [15:0] d;
  output [15:0] qtop, qnext;
  reg [15:0] q [0:7];

  assign qtop = q[0];
  assign qnext = q[1];
   
  always @(posedge clk or negedge reset)
    if(!reset) q[0] <= 0;
    else if(load) q[0] <= d;
    else if(pop) q[0] <= q[1];

  integer i;
  always @(posedge clk or negedge reset)
    for(i=1;i< 7;i=i+1)
      if(!reset) q[i] <= 0;
      else if(push) q[i] <= q[i-1];
      else if(pop) q[i] <= q[i+1];
   
  always @(posedge clk or negedge reset)
    if(!reset) q[N-1] <= 0;
    else if(push) q[N-1] <= q[N-2];

endmodule