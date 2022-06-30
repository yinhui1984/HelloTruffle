
//开发框架导入合约
const Calculator = artifacts.require("Calculator");

//接下来，我们定义用于测试的合约，然后将账户作为包含所有地址的参数传递。
contract("Calculator", accounts => {
    //it包含对我们要运行的测试的简短描述，
    // 它是一个包含所有测试相关脚本的异步函数
    it("should add two numbers", async () => {
        //cal: 定义存储已部署合约的实例。
        const cal = await Calculator.deployed();
        const result = await cal.add(4, 2);
        assert.equal(result.toNumber(), 6);
    }).timeout(10000);

    it("should subtract two numbers", async () => {
        const cal = await Calculator.deployed();
        const result = await cal.subtract(3, 2);
        assert.equal(result.toNumber(), 1);
    }).timeout(10000);
});