wordnet_go
====

## 使い方
https://github.com/katryo/wordnet_python と基本的に同じです。

まず wnjpn.db を http://nlpwww.nict.go.jp/wn-ja/ からダウンロードして、wn.goと同じディレクトリに入れます。

```
$ go run wn.go りんご
```

を実行すると、

```
o run wn.go りんご
     りんご
       edible_fruit
         green_goods
           food
             substance
               matter
                 concern
                   interest
                     curiosity
                       state_of_mind
                         temporary_state
                           state
                             administrative_division
                               district
                                 region
                                   location
                                     physical_object
                                       physical_entity
                                     activity
                                       human_action
                                         event
                                           psychological_feature
                                             abstract_entity
2014/06/16 23:56:06 おしまい
exit status 1
```

という出力になります。

- event
- activity
- abstract_entity
- state
- physical_object

あたりの抽象的な単語は、循環的なis-aの関係にあるので、無限にis-aで辿れます。が、このコードでは20までのis-aを辿ります。