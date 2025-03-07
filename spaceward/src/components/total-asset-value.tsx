import useWardenWardenV1Beta2 from "@/hooks/useWardenWardenV1Beta2";
import { WalletType } from "warden-protocol-wardenprotocol-client-ts/lib/warden.warden.v1beta2/rest";
import { ethers, formatEther } from "ethers";
import { useSpaceId } from "@/hooks/useSpaceId";
import { useCurrency } from "@/hooks/useCurrency";
import { useQueries } from "@tanstack/react-query";
import { Skeleton } from "./ui/skeleton";

const url = "https://rpc2.sepolia.org";
const chainId = 11155111;
const provider = new ethers.JsonRpcProvider(url);

async function getEthBalance(address: string) {
	const balance = await provider.getBalance(address);
	return balance;
}

const USDollar = new Intl.NumberFormat("en-US", {
	style: "currency",
	currency: "USD",
});
const Euro = new Intl.NumberFormat("en-US", {
	style: "currency",
	currency: "EUR",
});
const GBP = new Intl.NumberFormat("en-US", {
	style: "currency",
	currency: "GBP",
});

function TotalAssetValue() {
	const { currency } = useCurrency();
	const { spaceId } = useSpaceId();

	const { QueryKeys } = useWardenWardenV1Beta2();
	const { data: keysData } = QueryKeys(
		{
			type: WalletType.WALLET_TYPE_ETH,
			space_id: spaceId,
		},
		{},
		10
	);

	const totalBalanceQuery = useQueries({
		queries: keysData
			? keysData.pages
					.flatMap((page) => page.keys)
					.map((key) => key?.wallets?.[0]?.address)
					.map((ethAddr) => ({
						queryKey: ["eth-balance", chainId, ethAddr],
						queryFn: () => getEthBalance(ethAddr!),
						enabled: !!ethAddr,
					}))
			: [],
	});

	const totalBalance = totalBalanceQuery.reduce(
		(partialSum, result) => partialSum + BigInt(result.data || 0),
		BigInt(0)
	);
	// if we want some indicator to show that some of the queries are still pending (=> we're showing a partial result):
	const isPending = totalBalanceQuery.some(
		(result) => result.status === "loading"
	);

	return (
		<div>
			<p className="text-muted-foreground">Total asset value</p>
			{isPending ? (
				<Skeleton className="h-10 w-60" />
			) : (
				<span className="text-4xl font-bold flex">
					{currency === "usd"
						? USDollar.format(formatEther(totalBalance) * 2940)
						: currency === "eur"
							? Euro.format(formatEther(totalBalance) * 2756)
							: GBP.format(formatEther(totalBalance) * 2358)}
				</span>
			)}
		</div>
	);
}

export default TotalAssetValue;
