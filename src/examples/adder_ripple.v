module adder_ripple(a, b, q);
input [3:0] a, b;
output [3:0] q;
wire [3:0] cout;

fulladd add0 (.Q(q[0]), .COUT(cout[0]), .A(a[0]), .B(b[0]), .CIN(1'b0));
fulladd add1 (.Q(q[1]), .COUT(cout[1]), .A(a[1]), .B(b[1]), .CIN(cout[0]));
fulladd add2 (.Q(q[2]), .COUT(cout[2]), .A(a[2]), .B(b[2]), .CIN(cout[1]));
fulladd add3 (.Q(q[3]), .COUT(cout[3]), .A(a[3]), .B(b[3]), .CIN(cout[2]));

// fulladd add0 (a[0], b[0], 1'b0, q[0], cout[0]);
// fulladd add1 (a[1], b[1], cout[0], q[1], cout[1]);
// fulladd add2 (a[2], b[2], cout[1], q[2], cout[2]);
// fulladd add3 (a[3], b[3], cout[2], q[3], cout[3]);

endmodule