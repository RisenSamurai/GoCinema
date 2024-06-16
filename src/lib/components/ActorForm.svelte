<script>

</script>



<form class="flex flex-col bg-cinema-secondary w-full" method="POST">
    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Initials</legend>
        <label class="text-cinema-text " for="name">First Name</label>
        <input class="p-1 rounded-lg mb-2" type="text" name="name" id="name">
        <label class="text-cinema-text" for="name">Last Name</label>
        <input class="p-1 rounded-lg mb-2" type="text" name="lastName" id="lastName">
    </fieldset>
    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Biography</legend>
        <label class="text-cinema-text" for="birthday">Day of Birth</label>
        <input class="p-1 rounded-lg mb-2" type="date" name="birthday" id="birthday">
        <label class="text-cinema-text" for="gender">Gender</label>
        <select class="p-1 rounded-lg w-1/2 mb-2" id="gender" name="gender">
            <option value="male">Male</option>
            <option value="female">Female</option>
        </select>
        <label class="text-cinema-text" for="pob">Place of Birth</label>
        <input class="p-1 rounded-lg mb-2" type="text" name="pob" id="pob">
    </fieldset>
</form>