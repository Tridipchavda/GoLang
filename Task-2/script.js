// Grabing InputField by Id for Changing Field and Evaluation
var inputField = document.getElementById("inputField");

// boolean to check that the evaluation is happened or not
var isEqualPressed = false;

// function that helps to make the expression and also solve the expression
function addValueToExpression(val) {
  console.log("adding value to expression");
  
  
  if(inputField.value.length > 20){
    inputField.style.fontSize = '24px';
  }else{
    inputField.style.fontSize = '30px';
  }
  // If '=' is the input function value we need to evaluate the expression
  if (val === "=") {
    isEqualPressed = true;
    var ans = eval(inputField.value);
    // checking if the answer is float or not , if float .. round it to dynamic no. of digits
    if( ans - parseInt(ans) !== 0) {
        inputField.value = ans.toFixed(5)==0 ? ans : ans.toFixed(5);
        return;
    }else{
        inputField.value = ans;
    }
    
  }
  // Command to clear the input field
  else if (val === "clear") {
    inputField.value = "";
    // Setting boolean False when cleared 
    isEqualPressed = false;
  }
  // Only remove the last entered digit from the expression
  else if (val === "del") {
    inputField.value = inputField.value.substring(
      0,
      inputField.value.length - 1
    );
  }
  // Enter the operators with spacing between them
  else if (typeof val === "string" && val != ".") {
    isEqualPressed = false; 
    inputField.value += " " + val + " ";
  }
  // Else we join the value or operators enter by User in the expression Input Field
  else {
    // If evaluation happened at that instant it should not let 
    // user enter Integer in the Result , User must enter operator only after '=' like +,-,*,/
    if(!isEqualPressed){
        inputField.value += val;
    }
  }
}
