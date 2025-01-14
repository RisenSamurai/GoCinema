import {fail} from "@sveltejs/kit";

export const actions = {
    default: async ({ request }) => {
        const { email, password, confirmPassword } = await request.json();

        if (!email || !password || !confirmPassword) {
            return fail(400, { error: 'All fields are required!' });
        }

        if (password !== confirmPassword) {
            return fail(400, { error: 'Passwords do not match!' });
        }

        // Process data (e.g., save to database, authenticate user, etc.)

        const response = await fetch("http://localhost:8080/auth/sign-in", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({email, password}),
        });

        const data = await response.json();

        if (data.success) {

        }

        return { success: true };
    },
};