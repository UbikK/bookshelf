<script lang=ts>
    import {GetAllBooks} from "../../wailsjs/go/main/App"
    import type { main } from "../../wailsjs/go/models";
    export let books: main.Book[]

    
</script>

<div>
    {#if books}
        <ul>
            {#each books as book}
                <li>{book.name}</li>
            {/each}
        </ul>
    {/if}
</div>