<script lang="ts">
    import { onMount } from "svelte";
    import { Plaid } from "$lib";

    let linkToken = "";
    let plaidLinkInstance: any = null;

    onMount(async () => {
        const res = await fetch("api/create-link-token");
        const data = await res.json();
        linkToken = data.link_token;

        if (linkToken) {
            plaidLinkInstance = Plaid.initializePlaidLink(linkToken);
        }
    });

    const openPlaidLink = () => {
        if (plaidLinkInstance) {
            plaidLinkInstance.open();
        }
    };
</script>

<button on:click={openPlaidLink}>Connect Bank Account</button>

<style>
    button {
        background-color: #e0e0e0;
        padding: 8px 16px;
        border: none;
        border-radius: 4px;
        cursor: default;
        transition: background-color 0.2s;
    }

    button:hover {
        background-color: #d0d0d0;
        cursor: pointer;
    }
</style>
