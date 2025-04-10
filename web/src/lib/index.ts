export const Plaid = {
    initializePlaidLink: (linkToken: string) => {
        console.log("here");
        if (!window.Plaid) {
            console.error("Plaid Link SDK not loaded");
            return null;
        }

        return window.Plaid.create({
            token: linkToken,
            onSuccess: (public_token, metadata) => {
                console.log("Plaid onSuccess:", public_token, metadata);
                // TODO: Handle success (send public token to backend)
            },
            onExit: (err, metadata) => {
                console.error("Plaid onExit:", err, metadata);
                // Handle exit (e.g., user canceled or an error occurred)
            }
        });
    }
};
