package data

var (
	InputDefStm = `
	def num1 = 13;
	def  total= 0;
	def a= 5321;
	`
	DefStmExpected = []struct {
		ExpectedIdentifier string
		ExpectedValue      int
	}{
		{"num1", 13},
		{"total", 0},
		{"a", 5321},
	}

	ReturnStm = `
	return 545;
	return 101232;
	return 0;
	`
	ExpectedReturnValue = []int{545, 101232, 0}
	Identifier          = `varName;`

	IntegerLit = "81;"

	PrefixExpression = []struct {
		Input      string
		Operator   string
		IntOperand interface{}
	}{
		{Input: "!7;", Operator: "!", IntOperand: 7},
		{Input: "-42;", Operator: "-", IntOperand: 42},
		{Input: "!false;", Operator: "!", IntOperand: false},
		{Input: "!true;", Operator: "!", IntOperand: true},
	}

	InfixExpression = []struct {
		Input    string
		Left     interface{}
		Operator string
		Right    any
	}{
		{
			Input: "12 + 5;", Left: 12, Operator: "+", Right: 5,
		}, {
			Input: "12 - 5;", Left: 12, Operator: "-", Right: 5,
		}, {
			Input: "12*5;", Left: 12, Operator: "*", Right: 5,
		}, {
			Input: "12/5;", Left: 12, Operator: "/", Right: 5,
		}, {
			Input: "12 < 5;", Left: 12, Operator: "<", Right: 5,
		}, {
			Input: "12 <= 5;", Left: 12, Operator: "<=", Right: 5,
		}, {
			Input: "12 > 5;", Left: 12, Operator: ">", Right: 5,
		}, {
			Input: "12 >= 5;", Left: 12, Operator: ">=", Right: 5,
		}, {
			Input: "12 == 5;", Left: 12, Operator: "==", Right: 5,
		}, {
			Input: "12 != 5;", Left: 12, Operator: "!=", Right: 5,
		}, {
			Input: "true==true;", Left: true, Operator: "==", Right: true,
		}, {
			Input: "true != false;", Left: true, Operator: "!=", Right: false,
		},
	}

	ForLoopTestSimple = `for(def i=0;i<=10;i++){
													cnt + 1;
													if (cnt==0){
														return cnt;
													}
											}`

	PrecedenceOrder = []struct {
		Input    string
		Expected string
	}{
		{"-var1 * var2;", "((-var1)*var2)"},
		{"!-var1;", "(!(-var1))"},
		{"var1 + var2 + var3;", "((var1+var2)+var3)"},
		{"var1 * var2 * var3;", "((var1*var2)*var3)"},
		{"var1*var2/var3;", "((var1*var2)/var3)"},
		{"var1 + var2/var3;", "(var1+(var2/var3))"},
		{"var1 + var2 * var3 + var4 / var5 - var6;", "(((var1+(var2*var3))+(var4/var5))-var6)"},
		{"5>4 == 3<=4;", "((5>4)==(3<=4))"},
		{"5>=4 != 15<7;", "((5>=4)!=(15<7))"},
		{"3 + 4*5 == 3*1 + 4*5;", "((3+(4*5))==((3*1)+(4*5)))"},
		{"true;", "true"},
		{"3<5 == false;", "((3<5)==false)"},
		{"1 +(7 + 0) +44;", "((1+(7+0))+44)"},
		{"(7 + 7)*2;", "((7+7)*2)"},
		{"-10/(2+3);", "((-10)/(2+3))"},
		{"!(false != true);", "(!(false!=true))"},
		{"factorial(5);", "factorial(5)"},
		{"1+pow(2*5)/4;", "(1+(pow((2*5))/4))"},
		{"777++;", "(777++)"},
		{"max(1,65,2*11,100/2,max(100,12*30))", "max(1,65,(2*11),(100/2),max(100,(12*30)))"},
	}

	IfExpression = "if(m>=n) {m+1;} else{n+1;}"
	FunctionExp  = "function test(pr1,pr2){pr1+pr2;}"
	FunctionExp1 = "function test(){pr1+pr2;}"
	FunctionExp2 = "function test(pr1,pr2){}"

	FunctionCall = "functionName(arg1,arg2);"

	AssignExp = struct {
		Input        string
		ExpectedVal  interface{}
		ExpectedIden string
	}{
		"varname = 4645;",
		4645,
		"varname",
	}

	Arrays     = "[1,12 - 8 ,7]"
	ArrayIndex = "nums[7-4]"

	ClassExp = `class helloworld {
		
		def var1= 444;
		def var2=777;

		constructor() {

		}

		function toString() {
			out("hello world");
		}
	}`
)
