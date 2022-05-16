package envRead_viper

import (
    "fmt"

    "github.com/spf13/viper"
)

func ViperExample() {
    viper.SetConfigFile(".env")
    viper.ReadInConfig()

    fmt.Println(viper.Get("PORT"))
}
