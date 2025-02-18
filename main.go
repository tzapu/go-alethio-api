package main

import (
	"context"
	"fmt"
	"log"

	"github.com/corpetty/go-alethio-api/alethio"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	var address = "0x3378EeAf39DFFB316A95F31F17910CBb21ace6bB"
	client, err := alethio.NewClient(
		alethio.Opts.URL("https://api.goerli.aleth.io/v1"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// accountDetails, err := client.Account.GetDetails(context.Background(), address)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(accountDetails)

	etherTransfers, err := client.Account.GetEtherTransfers(context.Background(), address)
	if err != nil {
		fmt.Print(err)
	}

	etherTransfers, err = client.EtherTransfers.Get(context.Background(), etherTransfers.Links.Next)
	if err != nil {
		fmt.Print(err)
	}

	etherTransfers, err = client.EtherTransfers.GetPrev(context.Background(), etherTransfers)
	if err != nil {
		fmt.Print(err)
	}
	spew.Dump(etherTransfers)

	// var contract = "0xB8c77482e45F1F44dE1745F52C74426C631bDD52"
	// client := alethio.NewClient(nil)
	// contractDetails, err := client.Account.GetContract(context.Background(), contract)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(contractDetails)

	// var address = "0x2942577508e060ea092c0CD7802ae42c1CEA2BAe"
	// client := alethio.NewClient(nil)
	// accountTransactions, err := client.Account.GetTransactions(context.Background(), address)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(accountTransactions)

	// var address = "0x2942577508e060ea092c0CD7802ae42c1CEA2BAe"
	// client := alethio.NewClient(nil)
	// contractMessages, err := client.Account.GetContractMessages(context.Background(), address)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(contractMessages)

	// var address = "0x2942577508e060ea092c0CD7802ae42c1CEA2BAe"
	// client := alethio.NewClient(nil)
	// etherTransfers, err := client.Account.GetEtherTransfers(context.Background(), address)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(etherTransfers)

	// var address = "0x2942577508e060ea092c0CD7802ae42c1CEA2BAe"
	// client := alethio.NewClient(nil)
	// tokenTransfers, err := client.Account.GetTokenTransfers(context.Background(), address)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(tokenTransfers)

	// client := alethio.NewClient(nil)
	// blocks, err := client.Blocks.GetAll(context.Background())
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(blocks)

	// client2 := alethio.NewClient(nil)
	// var blockhash = "0x05245795a1b3ed9486387102539b55f71527f7cbe9afd3918642ed46735c30f0"
	// block, err := client2.Blocks.GetByHash(context.Background(), blockhash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(block)

	// client2 := alethio.NewClient(nil)
	// var blockNumber = 43243
	// block, err := client2.Blocks.GetByNumber(context.Background(), blockNumber)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(block)

	// client2 := alethio.NewClient(nil)
	// var label = "latest"
	// block, err := client2.Blocks.GetByLabel(context.Background(), label)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(block)

	// client2 := alethio.NewClient(nil)
	// var blockhash = "0x05245795a1b3ed9486387102539b55f71527f7cbe9afd3918642ed46735c30f0"
	// block, err := client2.Blocks.GetParent(context.Background(), blockhash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(block)

	// client2 := alethio.NewClient(nil)
	// var blockhash = "0x05245795a1b3ed9486387102539b55f71527f7cbe9afd3918642ed46735c30f0"
	// account, err := client2.Blocks.GetBeneficiary(context.Background(), blockhash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(account)

	// client2 := alethio.NewClient(nil)
	// var blockHash = "0x05245795a1b3ed9486387102539b55f71527f7cbe9afd3918642ed46735c30f0"
	// accountTransactions, err := client2.Blocks.GetTransactions(context.Background(), blockHash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(accountTransactions)

	// client2 := alethio.NewClient(nil)
	// var blockHash = "0x05245795a1b3ed9486387102539b55f71527f7cbe9afd3918642ed46735c30f0"
	// contractMessages, err := client2.Blocks.GetContractMessages(context.Background(), blockHash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(contractMessages)

	// client2 := alethio.NewClient(nil)
	// var blockHash = "0x05245795a1b3ed9486387102539b55f71527f7cbe9afd3918642ed46735c30f0"
	// logEntries, err := client2.Blocks.GetLogEntries(context.Background(), blockHash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(logEntries)

	// client2 := alethio.NewClient(nil)
	// var blockHash = "0x05245795a1b3ed9486387102539b55f71527f7cbe9afd3918642ed46735c30f0"
	// etherTransfers, err := client2.Blocks.GetEtherTransfers(context.Background(), blockHash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(etherTransfers)

	// client2 := alethio.NewClient(nil)
	// var blockHash = "0x05245795a1b3ed9486387102539b55f71527f7cbe9afd3918642ed46735c30f0"
	// tokenTransfers, err := client2.Blocks.GetTokenTransfers(context.Background(), blockHash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(tokenTransfers)

	// var address = "0xaA3fBFAf03cD50E6a44d27D10eB14333d1C02E52"
	// client := alethio.NewClient(nil)
	// contractDetails, err := client.Contracts.GetDetails(context.Background(), address)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(contractDetails)

	// var address = "0xaA3fBFAf03cD50E6a44d27D10eB14333d1C02E52"
	// client := alethio.NewClient(nil)
	// contractAccount, err := client.Contracts.GetAccount(context.Background(), address)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(contractAccount)

	// var address = "0xb8c77482e45f1f44de1745f52c74426c631bdd52"
	// client := alethio.NewClient(nil)
	// block, err := client.Contracts.GetCreationBlock(context.Background(), address)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(block)

	// var address = "0xb8c77482e45f1f44de1745f52c74426c631bdd52"
	// client := alethio.NewClient(nil)
	// transaction, err := client.Contracts.GetCreationTransaction(context.Background(), address)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(transaction)

	// var address = "0xb8c77482e45f1f44de1745f52c74426c631bdd52"
	// client := alethio.NewClient(nil)
	// message, err := client.Contracts.GetCreationMessage(context.Background(), address)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(message)

	// var address = "0xb8c77482e45f1f44de1745f52c74426c631bdd52"
	// client := alethio.NewClient(nil)
	// transactions, err := client.Contracts.GetTransactions(context.Background(), address)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(transactions)

	//var address = "0xb8c77482e45f1f44de1745f52c74426c631bdd52"
	//client, err := alethio.NewClient(nil)
	//contractMessages, err := client.Contracts.GetMessages(context.Background(), address)
	//if err != nil {
	//	fmt.Print(err)
	//}
	//spew.Dump(contractMessages)

	// var address = "0xb8c77482e45f1f44de1745f52c74426c631bdd52"
	// client := alethio.NewClient(nil)
	// logEntries, err := client.Contracts.GetLogEntries(context.Background(), address)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(logEntries)

	// var txHash = "0x8daaf1ffc66904e0d86e062c2d3b44d575f41203f9385d7a97d6ba41326b80aa"
	// client := alethio.NewClient(nil)
	// transaction, err := client.GetTransactionDetails(txHash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(transaction)

	// var txHash = "0x8daaf1ffc66904e0d86e062c2d3b44d575f41203f9385d7a97d6ba41326b80aa"
	// client := alethio.NewClient(nil)
	// sender, err := client.GetTransactionSender(txHash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(sender)

	// var txHash = "0x8daaf1ffc66904e0d86e062c2d3b44d575f41203f9385d7a97d6ba41326b80aa"
	// client := alethio.NewClient(nil)
	// to, err := client.GetTransactionDestination(txHash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(to)

	// var txHash = "0x8daaf1ffc66904e0d86e062c2d3b44d575f41203f9385d7a97d6ba41326b80aa"
	// client := alethio.NewClient(nil)
	// block, err := client.GetTransactionBlock(txHash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(block)

	// var txHash = "0x436fc7d21ed4a0a634f41b50ccb42fca12be7322de5bf9a20c97bdccbb5b2a04"
	// client := alethio.NewClient(nil)
	// contracts, err := client.GetTransactionCreatedContracts(txHash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(contracts)

	// var txHash = "0x436fc7d21ed4a0a634f41b50ccb42fca12be7322de5bf9a20c97bdccbb5b2a04"
	// client := alethio.NewClient(nil)
	// contractMessages, err := client.GetTransactionContractMessages(txHash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(contractMessages)

	// var id = "msg:0xb53fdc2d618b5c35da706b399f9201721a74082fdb46ee90e487c82e422a4c8c:4"
	// client := alethio.NewClient(nil)
	// contractMessages, err := client.GetContractMessages(id)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(contractMessages)
}
