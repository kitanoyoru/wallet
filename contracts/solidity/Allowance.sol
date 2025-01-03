// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../@openzeppelin/contracts/access/Ownable.sol";

contract Allowance is Ownable {

  event AllowanceChanged(
    address indexed beneficiary,
    address indexed sender,
    uint256 prevAmount,
    uint256 newAmount
  );

  mapping(address => uint256) public allowances;

  constructor() Ownable(msg.sender) {}

  function setAllowance(address _beneficiary, uint256 _amount) public onlyOwner {
    emit AllowanceChanged(
      _beneficiary,
      msg.sender,
      allowances[_beneficiary],
      _amount
    );
    allowances[_beneficiary] = _amount;
  }

  function increaseAllowance(address _beneficiary, uint256 _amount) public onlyOwner {
    uint256 newAmount = allowances[_beneficiary] + _amount;
    emit AllowanceChanged(
      _beneficiary,
      msg.sender,
      allowances[_beneficiary],
      newAmount
    );
    allowances[_beneficiary] = newAmount;
  }

  function reduceAllowance(address _beneficiary, uint256 _amount) public onlyOwner {
    uint256 newAmount = allowances[_beneficiary] - _amount;
    emit AllowanceChanged(
      _beneficiary,
      msg.sender,
      allowances[_beneficiary],
      newAmount
    );
    allowances[_beneficiary] = newAmount;
  }

  function isOwner() internal view returns (bool) {
    return owner() == msg.sender;
  }
}

