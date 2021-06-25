package dictionary

import (
	"fmt"
	"github.com/amjadjibon/spooky/pkg/constant"
)

func WordMeaningDecorator(responseModel []ResponseModel)  {
	for _, response := range responseModel {
		for _, meaning := range response.Meanings{
			for _, def := range meaning.Definitions{
				fmt.Printf("%s[%s]", constant.Red, responseModel[0].Word)
				fmt.Printf("%s - ", constant.White)
				fmt.Printf("%s(%s) ", constant.Purple, meaning.PartOfSpeech)
				fmt.Printf("%s%s\n",constant.Yellow, def.Definition)

				if len(def.Example) > 0 {
					fmt.Printf("%sExample: ", constant.Red)
					fmt.Printf("\n%s		%s\n",constant.Yellow, def.Example)
				}

				if len(def.Synonyms) > 0 {
					fmt.Printf("%sSynonyms: ", constant.Red)
					var syn string
					for i, synonym := range def.Synonyms {
						syn += fmt.Sprintf("\n		%d. %s",i+1, synonym)
					}
					fmt.Printf("%s%s\n\n",constant.Yellow, syn[:len(syn)-2])
				}
			}
		}
	}
}
