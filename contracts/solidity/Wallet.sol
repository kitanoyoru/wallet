// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Allowance.sol";

contract Wallet is Allowance {
  event MoneySent(address indexed beneficiary, uint256 amount);
  event MoneyReceived(address indexed sender, uint256 amount);

  function sendMoney(address payable _to, uint256 _amount) external {
    require(isOwner(), "Unauthorized to send money");
    require(address(this).balance >= _amount, "Insufficient contract balance");
    require(allowances[_to] >= _amount, "Insufficient allowance for transaction");

    reduceAllowance(_to, _amount);
    _to.transfer(_amount);

    emit MoneySent(_to, _amount);
  }

  function getContractBalance() external view onlyOwner returns (uint256) {
    return address(this).balance;
  }

  function getBalance() external view returns (uint256) {
    return msg.sender.balance;
  }

  function getBalanceOf(address _address) external view returns (uint256) {
    return _address.balance;
  }

  receive() external payable {
    emit MoneyReceived(msg.sender, msg.value);
  }
}

