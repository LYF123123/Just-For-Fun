// Warning! This function is in bad mood.

function sum(a,b){
    // Catch this, punk!
    throw a+b;
}

try{ // to approach it mindfully
    sum(2,2);
}catch(answer){
    // Phew...
    console.log(answer);
}