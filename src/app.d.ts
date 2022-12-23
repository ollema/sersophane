import PocketBase from 'pocketbase';

declare namespace App {
	interface Locals {
		pb: PocketBase;
		user?: {
			id: string;
			email: string;
			username: string;
			name: string;
			avatar: string;
		};
	}
}
