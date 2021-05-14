module adder_test;

reg [3:0] a, b;
wire [3:0] q;

adder adder(a,b,q);

initial begin
       a = 0; b = 0;
       #100 a = 1; b = 0;
       #100 a = 2; b = 1;
       #100 a = 5; b = 7;
end

initial begin
       $monitor("a=%d, b=%d, q=%d", a, b, q);
       $dumpfile("adder.vcd");
       $dumpvars(0, adder_test);
end

endmodule