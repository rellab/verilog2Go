module elelock_test;

reg clk, reset, close;
reg [9: 0] tenkey;
wire lock;

elelock elelock(clk, reset, tenkey, close, lock);

always #5  clk = ~clk;

initial begin
       clk = 0;reset = 0; 
       #150     $finish;
end

initial begin
       $monitor("CLK=%d, RES=%d, Q=%d", clk, reset, lock);
       $dumpfile("elelock.vcd");
       $dumpvars(0, elelock_test);
end

endmodule