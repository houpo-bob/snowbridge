import "dotenv/config"
import { historyV2, subsquid } from "@snowbridge/api"

const monitor = async () => {
    const graphqlApiUrl = process.env["GRAPHQL_API_URL"]
    if (!graphqlApiUrl) throw Error("'GRAPHQL_API_URL' env var not set.")
    console.log("To Ethereum transfers:")
    const toEthereums = await historyV2.toEthereumHistory(graphqlApiUrl)
    console.log(JSON.stringify(toEthereums, null, 2))

    console.log("To Polkadot transfers:")
    const toPolkadots = await historyV2.toPolkadotHistory(graphqlApiUrl)
    console.log(JSON.stringify(toPolkadots, null, 2))

    console.log("All transfers order by time:")
    const transfers = [...toEthereums, ...toPolkadots]
    transfers.sort((a, b) => b.info.when.getTime() - a.info.when.getTime())
    console.log(JSON.stringify(transfers, null, 2))

    console.log("To Polkadot transfer by id:")
    const toPolkadot = await historyV2.toPolkadotTransferById(
        graphqlApiUrl,
        "0xb56662848712da9769a2122ca0d24d199ef7af7c8aedee43778dadbe1c42ebc6"
    )
    console.log(JSON.stringify(toPolkadot, null, 2))

    console.log("To Ethereum transfer by id:")
    const toEthereum = await historyV2.toEthereumTransferById(
        graphqlApiUrl,
        "0x04b7a6c7552d2890094dfe43e037cb5f5495fec2419f33b0072439a9ee7629a0"
    )
    console.log(JSON.stringify(toEthereum, null, 2))

    const estimatedDeliveryTime = await subsquid.fetchEstimatedDeliveryTime(
        graphqlApiUrl,
        "0xc173fac324158e77fb5840738a1a541f633cbec8884c6a601c567d2b376a0539"
    )
    console.log(estimatedDeliveryTime)
    const latestBlock = await subsquid.fetchLatestBlocksSynced(graphqlApiUrl, true)
    console.log(latestBlock)
}

monitor()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("Error:", error)
        process.exit(1)
    })
