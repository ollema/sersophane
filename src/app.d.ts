import PocketBase from 'pocketbase';

declare global {
	namespace App {
		interface Locals {
			pb: PocketBase;
			user?: {
				id: string;
				username: string;
				name: string;
				email: string;
				avatar: string;
			};
		}
	}
}
