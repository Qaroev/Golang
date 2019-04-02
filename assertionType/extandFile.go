package main

/*
public class Program{

    public static void main(String[] args) {

        Account<String> acc1 = new Account<String>("1876", 4500);
        Account<String> acc2 = new Account<String>("3476", 1500);

        Transaction<Account<String>> tran1 = new Transaction<Account<String>>(acc1,acc2, 4000);
        tran1.execute();
        tran1 = new Transaction<Account<String>>(acc1,acc2, 4000);
        tran1.execute();
    }
}
class Transaction<T extends Account<String>>{

    private T from;     // с какого счета перевод
    private T to;       // на какой счет перевод
    private int sum;    // сумма перевода

    Transaction(T from, T to, int sum){
        this.from = from;
        this.to = to;
        this.sum = sum;
    }
    public void execute(){

        if (from.getSum() > sum)
        {
            from.setSum(from.getSum() - sum);
            to.setSum(to.getSum() + sum);
            System.out.printf("Account %s: %d \nAccount %s: %d \n",
                from.getId(), from.getSum(),to.getId(), to.getSum());
        }
        else{
            System.out.printf("Operation is invalid");
        }
    }
}
class Account<T>{

    private T id;
    private int sum;

    Account(T id, int sum){
        this.id = id;
        this.sum = sum;
    }

    public T getId() { return id; }
    public int getSum() { return sum; }
    public void setSum(int sum) { this.sum = sum; }
}
 */
import "fmt"

type ExtanAccounts struct {
	id  string
	sum int
}

func NewExtanAccounts(id string, sum int) ExtanAccounts {
	a := ExtanAccounts{}
	a.id = id
	a.sum = sum
	return a
}

func (a ExtanAccounts) GetId() string {
	return a.id
}

func (a ExtanAccounts) GetSum() int {
	return a.sum
}

func (a ExtanAccounts) SetSum(sum int) {
	a.sum = sum
}

type Transaction struct {
	from ExtanAccounts
	To   ExtanAccounts
	sum  int
}

func NewTransaction(from ExtanAccounts, To ExtanAccounts, sum int) Transaction {
	t := Transaction{}
	t.from = from
	t.To = To
	t.sum = sum
	return t
}

func (t Transaction) execute() {
	if t.from.GetSum() > t.sum {
		t.from.SetSum(t.from.GetSum() - t.sum)
		t.To.SetSum(t.To.GetSum() + t.sum)
		fmt.Println(fmt.Sprintf("Account %v: %v \nAccount %v: %v \n",
			t.from.GetId(), t.from.GetSum(), t.To.GetId(), t.To.GetSum()))
	} else {
		fmt.Println("Operation is invalid")
	}
}

func main() {
	acc := NewExtanAccounts("1876", 4500)
	acc2 := NewExtanAccounts("3476", 1500)

	tran := NewTransaction(acc, acc2, 4000)
	tran.execute()
	tran = NewTransaction(acc, acc2, 4000)
	tran.execute()
}
