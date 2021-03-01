module counter_unit(ck, res, q);
input ck, res;
output [3:0] q;
wire [3:0] ca;

cnt_unit cu0(.CK(ck), .RES(res), .EN(1'b1), .Q(q[0]), .CA(ca[0]));
cnt_unit cu1(.CK(ck), .RES(res), .EN(ca[0]), .Q(q[1]), .CA(ca[1]));
cnt_unit cu2(.CK(ck), .RES(res), .EN(ca[1]), .Q(q[2]), .CA(ca[2]));
cnt_unit cu3(.CK(ck), .RES(res), .EN(ca[2]), .Q(q[3]), .CA(ca[3]));

// cnt_unit cu0(ck, res, 1'b1, q[0], ca[0]);
// cnt_unit cu1(ck, res, ca[0], q[1], ca[1]);
// cnt_unit cu2(ck, res, ca[1], q[2], ca[2]);
// cnt_unit cu3(ck, res, ca[2], q[3], ca[3]);

endmodule