<script lang=ts>
	import { createEventDispatcher } from "svelte";
	import { CreateBook } from "../../wailsjs/go/main/App";

    let name: string;
    let id: number
    const dispatch = createEventDispatcher();
    
    function createBook(name: string) {
        CreateBook(name).then(newId =>{ 
            id = newId
            dispatch('newbook')
        })
    }
</script>

<div>
    <form method="post">
        <label for="name">Name</label>
        <input bind:value={name} class="input" id="name" type="text"/>
        <button type="button" on:click|preventDefault={() => createBook(name)}>Submit</button>
        
    </form>
</div>