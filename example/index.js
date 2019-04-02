class Person {
    constructor(inn, name) {
        this.name = name
        this.inn = inn
    }
}

class SecretAgent extends Person {
    constructor(LicenseToKill, inn, name) {
        super(inn, name);
        this.LicenseToKill = LicenseToKill
    }

    GetName() {
        return this.LicenseToKill
    }
}


var s = new SecretAgent(true, "Hello", "world");

var res = s.GetName();

console.log(res)