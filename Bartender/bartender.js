var your_drink="coke";

var reverse = function(s){
    return s.split("").reverse().join("");
};

var bartender = {
    str1:"ers",
    str2:reverse("rap"),
    str3:"amet",
    request:function(preference){
        return preference+".Secret word: "+this.str2+this.str3+this.str1;
    }
};

var res=bartender.request(your_drink);
console.log(res);