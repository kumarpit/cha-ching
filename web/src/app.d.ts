declare global {
    interface Plaid {
        create: (config: PlaidLinkConfiguration) => any;
    }

    interface PlaidLinkConfiguration {
        token: string;
        onSuccess: (public_token: string, metadata: any) => void;
        onExit: (err: any, metadata: any) => void;
    }

    interface Window {
        Plaid: Plaid;
    }
    namespace App {
        // interface Error {}
        // interface Locals {}
        // interface PageData {}
        // interface PageState {}
        // interface Platform {}
    }
}

export { };
