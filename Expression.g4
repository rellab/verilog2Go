grammar Expression;

start
	: expression EOF
	;

expression returns [int number]
    : op=('!'|'+'|'-'|'~'|'|'|'&') expression {$number = 1;}
    | expression op='**' expression { $number = 2; }
    | expression op=('*'|'/'|'%') expression { $number = 3; }
    | expression op=('+'|'-') expression { $number = 4; }
    | expression op=('<<'|'>>'|'<<<'|'>>>') expression { $number = 5; }
    | expression op=('<'|'<='|'>'|'>=') expression { $number = 6; }
    | expression op=('=='|'!='|'==='|'!==') expression { $number = 7; }
    | expression op=('&'|'~&') expression { $number = 8; }
    | expression op=('^'|'~^') expression { $number = 9; }
    | expression op=('|'|'~|') expression { $number = 10; }
    | expression op='&&' expression { $number = 11; }
    | expression op='||' expression { $number = 12; }
    | id=Number { $number = 17; }
    | id=ID { $number = 13; }
    |	'(' expression ')' { $number = 14; }
    | expression '(' expression ')' { $number = 15; }
    | ID'['Decimal ']' { $number = 16; }
    ;

Number
   : Decimal_number
   | Binary_number
   | Octal_number
   | Hex_number
   ;

ID
   : [a-zA-Z_] [a-zA-Z0-9$_']*
   | Binary_number
   | Octal_number
   | Hex_number
   ;
Decimal_number
   : (Size)? Decimal_base Unsigned_number
   ;
Binary_number
   : (Size)? Binary_base Binary_value
   ;
Octal_number
   : (Size)? Octal_base Octal_value
   ;
Hex_number
   : (Size)? Hex_base Hex_value
   ;
fragment Size
   : Non_zero_unsigned_number
   ;
fragment Non_zero_unsigned_number
   : Non_zero_decimal_digit ('_' | Decimal_digit)*
   ;
fragment Unsigned_number
   : Decimal_digit ('_' | Decimal_digit)*
   ;
fragment Binary_value
   : Binary_digit ('_' | Binary_digit)*
   ;
fragment Octal_value
   : Octal_digit ('_' | Octal_digit)*
   ;
fragment Hex_value
   : Hex_digit ('_' | Hex_digit)*
   ;
fragment Decimal_base
   : '\'' [sS]? [dD]
   ;


fragment Binary_base
   : '\'' [sS]? [bB]
   ;


fragment Octal_base
   : '\'' [sS]? [oO]
   ;


fragment Hex_base
   : '\'' [sS]? [hH]
   ;


fragment Non_zero_decimal_digit
   : [1-9]
   ;


fragment Decimal_digit
   : [0-9]
   ;

Decimal
	: [0-9]
	;


fragment Binary_digit
   : X_digit | Z_digit | [01]
   ;


fragment Octal_digit
   : X_digit | Z_digit | [0-7]
   ;


fragment Hex_digit
   : X_digit | Z_digit | [0-9a-fA-F]
   ;


fragment X_digit
   : [xX]
   ;


fragment Z_digit
   : [zZ?]
   ;

White_space
   : [ \t\n\r] + -> channel (HIDDEN)
   ;