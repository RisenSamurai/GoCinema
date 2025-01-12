

<script>
    let email = $state();
    let password = $state();
    let confirmPassword = $state();
    let message = $state('');

    async function handleSubmit(event) {
        event.preventDefault();
        
        if(password === confirmPassword) {
            const response = await fetch('/sign-up', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ email, password}),
            })

            const result = await response.json();
            if(result.success){
                console.log(result.data);
            } else {
                console.log("Error occurred");
            }
        } else {
            message = 'Passwords do not match';
        }



    }


</script>



<div class="flex flex-col bg-cinema-secondary justify-center h-screen text-white self-center">

    {#if message}<span>{message}</span>{/if}
    <form onsubmit={handleSubmit} class="flex flex-col justify-center self-center">
        <label for="email">Email</label>
        <input class="text-black" type="email" id="email" name="email" bind:value={email} required/>
        <label for="password">Password</label>
        <input class="text-black" type="password" id="password" name="password" bind:value={password} required/>
        <label for="password_repeat">Confirm Password</label>
        <input class="text-black" type="password" id="password_repeat" name="password_repeat" bind:value={confirmPassword} required/>

        <button type="submit" class="btn btn-primary">Submit</button>
    </form>

    <span class="flex self-center">OR</span>

    <div class="flex self-center">
        <span>Google</span>
    </div>

</div>