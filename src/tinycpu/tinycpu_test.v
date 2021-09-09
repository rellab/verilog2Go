module tinycpu_test;

  reg clk,reset,run;
  reg [15:0] in;
  wire [2:0] cs;
  wire [15:0] irout, qtop, dbus, out;
  wire [11:0] pcout, abus;

  tinycpu tinycpu(clk, reset, run, in, cs, pcout, irout, qtop, abus, dbus, out);

  always #5  clk = ~clk;

  initial begin
    clk = 0; reset = 0; run = 0; in = 0;
    #50 run = 1;
    #10 in = 2;
    #100 in = 5;
    #500 $finish;
  end

  initial begin
        $monitor("clk=%d, reset=%d, run=%d", clk, reset, run);
        $dumpfile("tinycpu.vcd");
        $dumpvars(0, tinycpu_test);
  end

endmodule