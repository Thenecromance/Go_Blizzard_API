package main

import (
	"Unofficial_API/wow"
	"context"
	"encoding/json"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	res := wow.CNPlayerSummary(context.Background(), "黑色卷卷毛", "blanchard")
	buff, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	os.WriteFile("out.json", buff, 0644)
}
