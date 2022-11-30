async function main() {
    const FlowShieldDao = await ethers.getContractFactory("FlowShieldDao");
    console.log("Deploying FlowShieldDao...");

    const contract = await upgrades.deployProxy(FlowShieldDao, { initializer: 'initialize', kind: "transparent",});
    await contract.deployed();
    console.log("FlowShieldDao deployed to:", contract.address);
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });