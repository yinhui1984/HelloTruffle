
const Web3 = require('web3');
const web3 = new Web3(new Web3.providers.HttpProvider("http://127.0.0.1:8545"));

//web3.eth.getAccounts().then(console.log)

//modify the following line to your own contract address
let contractAddress = "0x6B62e4E253823FBC65E0B93d63ee149350158a18";

//copy abi from ./build/contracts/Calculator.json
let abi = [
    {
        "inputs": [
            {
                "internalType": "int256",
                "name": "a",
                "type": "int256"
            },
            {
                "internalType": "int256",
                "name": "b",
                "type": "int256"
            }
        ],
        "name": "add",
        "outputs": [
            {
                "internalType": "int256",
                "name": "",
                "type": "int256"
            }
        ],
        "stateMutability": "pure",
        "type": "function",
        "constant": true
    },
    {
        "inputs": [
            {
                "internalType": "int256",
                "name": "a",
                "type": "int256"
            },
            {
                "internalType": "int256",
                "name": "b",
                "type": "int256"
            }
        ],
        "name": "subtract",
        "outputs": [
            {
                "internalType": "int256",
                "name": "",
                "type": "int256"
            }
        ],
        "stateMutability": "pure",
        "type": "function",
        "constant": true
    }
]
let contract = new web3.eth.Contract(abi,contractAddress);

contract.methods.add(4,2).call().then(console.log);
contract.methods.subtract(3,2).call().then(console.log);

