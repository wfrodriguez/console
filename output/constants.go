package output

// type Color string
// type Text string

var showColor = true

// Base Constants

// Text Constants
const (
	Bold      Text = "1"
	Dim       Text = "2"
	Italic    Text = "3"
	Underline Text = "4"
	Blink     Text = "5"
	Reverse   Text = "7"
	Hidden    Text = "8"
	Strike    Text = "9"
)

// // 256 Xterm colours
// const ( // {{{
// 	Colour0   Color = "0"
// 	Colour1   Color = "1"
// 	Colour2   Color = "2"
// 	Colour3   Color = "3"
// 	Colour4   Color = "4"
// 	Colour5   Color = "5"
// 	Colour6   Color = "6"
// 	Colour7   Color = "7"
// 	Colour8   Color = "8"
// 	Colour9   Color = "9"
// 	Colour10  Color = "10"
// 	Colour11  Color = "11"
// 	Colour12  Color = "12"
// 	Colour13  Color = "13"
// 	Colour14  Color = "14"
// 	Colour15  Color = "15"
// 	Colour16  Color = "16"
// 	Colour17  Color = "17"
// 	Colour18  Color = "18"
// 	Colour19  Color = "19"
// 	Colour20  Color = "20"
// 	Colour21  Color = "21"
// 	Colour22  Color = "22"
// 	Colour23  Color = "23"
// 	Colour24  Color = "24"
// 	Colour25  Color = "25"
// 	Colour26  Color = "26"
// 	Colour27  Color = "27"
// 	Colour28  Color = "28"
// 	Colour29  Color = "29"
// 	Colour30  Color = "30"
// 	Colour31  Color = "31"
// 	Colour32  Color = "32"
// 	Colour33  Color = "33"
// 	Colour34  Color = "34"
// 	Colour35  Color = "35"
// 	Colour36  Color = "36"
// 	Colour37  Color = "37"
// 	Colour38  Color = "38"
// 	Colour39  Color = "39"
// 	Colour40  Color = "40"
// 	Colour41  Color = "41"
// 	Colour42  Color = "42"
// 	Colour43  Color = "43"
// 	Colour44  Color = "44"
// 	Colour45  Color = "45"
// 	Colour46  Color = "46"
// 	Colour47  Color = "47"
// 	Colour48  Color = "48"
// 	Colour49  Color = "49"
// 	Colour50  Color = "50"
// 	Colour51  Color = "51"
// 	Colour52  Color = "52"
// 	Colour53  Color = "53"
// 	Colour54  Color = "54"
// 	Colour55  Color = "55"
// 	Colour56  Color = "56"
// 	Colour57  Color = "57"
// 	Colour58  Color = "58"
// 	Colour59  Color = "59"
// 	Colour60  Color = "60"
// 	Colour61  Color = "61"
// 	Colour62  Color = "62"
// 	Colour63  Color = "63"
// 	Colour64  Color = "64"
// 	Colour65  Color = "65"
// 	Colour66  Color = "66"
// 	Colour67  Color = "67"
// 	Colour68  Color = "68"
// 	Colour69  Color = "69"
// 	Colour70  Color = "70"
// 	Colour71  Color = "71"
// 	Colour72  Color = "72"
// 	Colour73  Color = "73"
// 	Colour74  Color = "74"
// 	Colour75  Color = "75"
// 	Colour76  Color = "76"
// 	Colour77  Color = "77"
// 	Colour78  Color = "78"
// 	Colour79  Color = "79"
// 	Colour80  Color = "80"
// 	Colour81  Color = "81"
// 	Colour82  Color = "82"
// 	Colour83  Color = "83"
// 	Colour84  Color = "84"
// 	Colour85  Color = "85"
// 	Colour86  Color = "86"
// 	Colour87  Color = "87"
// 	Colour88  Color = "88"
// 	Colour89  Color = "89"
// 	Colour90  Color = "90"
// 	Colour91  Color = "91"
// 	Colour92  Color = "92"
// 	Colour93  Color = "93"
// 	Colour94  Color = "94"
// 	Colour95  Color = "95"
// 	Colour96  Color = "96"
// 	Colour97  Color = "97"
// 	Colour98  Color = "98"
// 	Colour99  Color = "99"
// 	Colour100 Color = "100"
// 	Colour101 Color = "101"
// 	Colour102 Color = "102"
// 	Colour103 Color = "103"
// 	Colour104 Color = "104"
// 	Colour105 Color = "105"
// 	Colour106 Color = "106"
// 	Colour107 Color = "107"
// 	Colour108 Color = "108"
// 	Colour109 Color = "109"
// 	Colour110 Color = "110"
// 	Colour111 Color = "111"
// 	Colour112 Color = "112"
// 	Colour113 Color = "113"
// 	Colour114 Color = "114"
// 	Colour115 Color = "115"
// 	Colour116 Color = "116"
// 	Colour117 Color = "117"
// 	Colour118 Color = "118"
// 	Colour119 Color = "119"
// 	Colour120 Color = "120"
// 	Colour121 Color = "121"
// 	Colour122 Color = "122"
// 	Colour123 Color = "123"
// 	Colour124 Color = "124"
// 	Colour125 Color = "125"
// 	Colour126 Color = "126"
// 	Colour127 Color = "127"
// 	Colour128 Color = "128"
// 	Colour129 Color = "129"
// 	Colour130 Color = "130"
// 	Colour131 Color = "131"
// 	Colour132 Color = "132"
// 	Colour133 Color = "133"
// 	Colour134 Color = "134"
// 	Colour135 Color = "135"
// 	Colour136 Color = "136"
// 	Colour137 Color = "137"
// 	Colour138 Color = "138"
// 	Colour139 Color = "139"
// 	Colour140 Color = "140"
// 	Colour141 Color = "141"
// 	Colour142 Color = "142"
// 	Colour143 Color = "143"
// 	Colour144 Color = "144"
// 	Colour145 Color = "145"
// 	Colour146 Color = "146"
// 	Colour147 Color = "147"
// 	Colour148 Color = "148"
// 	Colour149 Color = "149"
// 	Colour150 Color = "150"
// 	Colour151 Color = "151"
// 	Colour152 Color = "152"
// 	Colour153 Color = "153"
// 	Colour154 Color = "154"
// 	Colour155 Color = "155"
// 	Colour156 Color = "156"
// 	Colour157 Color = "157"
// 	Colour158 Color = "158"
// 	Colour159 Color = "159"
// 	Colour160 Color = "160"
// 	Colour161 Color = "161"
// 	Colour162 Color = "162"
// 	Colour163 Color = "163"
// 	Colour164 Color = "164"
// 	Colour165 Color = "165"
// 	Colour166 Color = "166"
// 	Colour167 Color = "167"
// 	Colour168 Color = "168"
// 	Colour169 Color = "169"
// 	Colour170 Color = "170"
// 	Colour171 Color = "171"
// 	Colour172 Color = "172"
// 	Colour173 Color = "173"
// 	Colour174 Color = "174"
// 	Colour175 Color = "175"
// 	Colour176 Color = "176"
// 	Colour177 Color = "177"
// 	Colour178 Color = "178"
// 	Colour179 Color = "179"
// 	Colour180 Color = "180"
// 	Colour181 Color = "181"
// 	Colour182 Color = "182"
// 	Colour183 Color = "183"
// 	Colour184 Color = "184"
// 	Colour185 Color = "185"
// 	Colour186 Color = "186"
// 	Colour187 Color = "187"
// 	Colour188 Color = "188"
// 	Colour189 Color = "189"
// 	Colour190 Color = "190"
// 	Colour191 Color = "191"
// 	Colour192 Color = "192"
// 	Colour193 Color = "193"
// 	Colour194 Color = "194"
// 	Colour195 Color = "195"
// 	Colour196 Color = "196"
// 	Colour197 Color = "197"
// 	Colour198 Color = "198"
// 	Colour199 Color = "199"
// 	Colour200 Color = "200"
// 	Colour201 Color = "201"
// 	Colour202 Color = "202"
// 	Colour203 Color = "203"
// 	Colour204 Color = "204"
// 	Colour205 Color = "205"
// 	Colour206 Color = "206"
// 	Colour207 Color = "207"
// 	Colour208 Color = "208"
// 	Colour209 Color = "209"
// 	Colour210 Color = "210"
// 	Colour211 Color = "211"
// 	Colour212 Color = "212"
// 	Colour213 Color = "213"
// 	Colour214 Color = "214"
// 	Colour215 Color = "215"
// 	Colour216 Color = "216"
// 	Colour217 Color = "217"
// 	Colour218 Color = "218"
// 	Colour219 Color = "219"
// 	Colour220 Color = "220"
// 	Colour221 Color = "221"
// 	Colour222 Color = "222"
// 	Colour223 Color = "223"
// 	Colour224 Color = "224"
// 	Colour225 Color = "225"
// 	Colour226 Color = "226"
// 	Colour227 Color = "227"
// 	Colour228 Color = "228"
// 	Colour229 Color = "229"
// 	Colour230 Color = "230"
// 	Colour231 Color = "231"
// 	Colour232 Color = "232"
// 	Colour233 Color = "233"
// 	Colour234 Color = "234"
// 	Colour235 Color = "235"
// 	Colour236 Color = "236"
// 	Colour237 Color = "237"
// 	Colour238 Color = "238"
// 	Colour239 Color = "239"
// 	Colour240 Color = "240"
// 	Colour241 Color = "241"
// 	Colour242 Color = "242"
// 	Colour243 Color = "243"
// 	Colour244 Color = "244"
// 	Colour245 Color = "245"
// 	Colour246 Color = "246"
// 	Colour247 Color = "247"
// 	Colour248 Color = "248"
// 	Colour249 Color = "249"
// 	Colour250 Color = "250"
// 	Colour251 Color = "251"
// 	Colour252 Color = "252"
// 	Colour253 Color = "253"
// 	Colour254 Color = "254"
// 	Colour255 Color = "255"

// 	// 256 Xterm colours Hexadecimal
// 	HexColour0   Color = "#000000"
// 	HexColour1   Color = "#800000"
// 	HexColour2   Color = "#008000"
// 	HexColour3   Color = "#808000"
// 	HexColour4   Color = "#000080"
// 	HexColour5   Color = "#800080"
// 	HexColour6   Color = "#008080"
// 	HexColour7   Color = "#c0c0c0"
// 	HexColour8   Color = "#808080"
// 	HexColour9   Color = "#ff0000"
// 	HexColour10  Color = "#00ff00"
// 	HexColour11  Color = "#ffff00"
// 	HexColour12  Color = "#0000ff"
// 	HexColour13  Color = "#ff00ff"
// 	HexColour14  Color = "#00ffff"
// 	HexColour15  Color = "#ffffff"
// 	HexColour16  Color = "#000000"
// 	HexColour17  Color = "#00005f"
// 	HexColour18  Color = "#000087"
// 	HexColour19  Color = "#0000af"
// 	HexColour20  Color = "#0000d7"
// 	HexColour21  Color = "#0000ff"
// 	HexColour22  Color = "#005f00"
// 	HexColour23  Color = "#005f5f"
// 	HexColour24  Color = "#005f87"
// 	HexColour25  Color = "#005faf"
// 	HexColour26  Color = "#005fd7"
// 	HexColour27  Color = "#005fff"
// 	HexColour28  Color = "#008700"
// 	HexColour29  Color = "#00875f"
// 	HexColour30  Color = "#008787"
// 	HexColour31  Color = "#0087af"
// 	HexColour32  Color = "#0087d7"
// 	HexColour33  Color = "#0087ff"
// 	HexColour34  Color = "#00af00"
// 	HexColour35  Color = "#00af5f"
// 	HexColour36  Color = "#00af87"
// 	HexColour37  Color = "#00afaf"
// 	HexColour38  Color = "#00afd7"
// 	HexColour39  Color = "#00afff"
// 	HexColour40  Color = "#00d700"
// 	HexColour41  Color = "#00d75f"
// 	HexColour42  Color = "#00d787"
// 	HexColour43  Color = "#00d7af"
// 	HexColour44  Color = "#00d7d7"
// 	HexColour45  Color = "#00d7ff"
// 	HexColour46  Color = "#00ff00"
// 	HexColour47  Color = "#00ff5f"
// 	HexColour48  Color = "#00ff87"
// 	HexColour49  Color = "#00ffaf"
// 	HexColour50  Color = "#00ffd7"
// 	HexColour51  Color = "#00ffff"
// 	HexColour52  Color = "#5f0000"
// 	HexColour53  Color = "#5f005f"
// 	HexColour54  Color = "#5f0087"
// 	HexColour55  Color = "#5f00af"
// 	HexColour56  Color = "#5f00d7"
// 	HexColour57  Color = "#5f00ff"
// 	HexColour58  Color = "#5f5f00"
// 	HexColour59  Color = "#5f5f5f"
// 	HexColour60  Color = "#5f5f87"
// 	HexColour61  Color = "#5f5faf"
// 	HexColour62  Color = "#5f5fd7"
// 	HexColour63  Color = "#5f5fff"
// 	HexColour64  Color = "#5f8700"
// 	HexColour65  Color = "#5f875f"
// 	HexColour66  Color = "#5f8787"
// 	HexColour67  Color = "#5f87af"
// 	HexColour68  Color = "#5f87d7"
// 	HexColour69  Color = "#5f87ff"
// 	HexColour70  Color = "#5faf00"
// 	HexColour71  Color = "#5faf5f"
// 	HexColour72  Color = "#5faf87"
// 	HexColour73  Color = "#5fafaf"
// 	HexColour74  Color = "#5fafd7"
// 	HexColour75  Color = "#5fafff"
// 	HexColour76  Color = "#5fd700"
// 	HexColour77  Color = "#5fd75f"
// 	HexColour78  Color = "#5fd787"
// 	HexColour79  Color = "#5fd7af"
// 	HexColour80  Color = "#5fd7d7"
// 	HexColour81  Color = "#5fd7ff"
// 	HexColour82  Color = "#5fff00"
// 	HexColour83  Color = "#5fff5f"
// 	HexColour84  Color = "#5fff87"
// 	HexColour85  Color = "#5fffaf"
// 	HexColour86  Color = "#5fffd7"
// 	HexColour87  Color = "#5fffff"
// 	HexColour88  Color = "#870000"
// 	HexColour89  Color = "#87005f"
// 	HexColour90  Color = "#870087"
// 	HexColour91  Color = "#8700af"
// 	HexColour92  Color = "#8700d7"
// 	HexColour93  Color = "#8700ff"
// 	HexColour94  Color = "#875f00"
// 	HexColour95  Color = "#875f5f"
// 	HexColour96  Color = "#875f87"
// 	HexColour97  Color = "#875faf"
// 	HexColour98  Color = "#875fd7"
// 	HexColour99  Color = "#875fff"
// 	HexColour100 Color = "#878700"
// 	HexColour101 Color = "#87875f"
// 	HexColour102 Color = "#878787"
// 	HexColour103 Color = "#8787af"
// 	HexColour104 Color = "#8787d7"
// 	HexColour105 Color = "#8787ff"
// 	HexColour106 Color = "#87af00"
// 	HexColour107 Color = "#87af5f"
// 	HexColour108 Color = "#87af87"
// 	HexColour109 Color = "#87afaf"
// 	HexColour110 Color = "#87afd7"
// 	HexColour111 Color = "#87afff"
// 	HexColour112 Color = "#87d700"
// 	HexColour113 Color = "#87d75f"
// 	HexColour114 Color = "#87d787"
// 	HexColour115 Color = "#87d7af"
// 	HexColour116 Color = "#87d7d7"
// 	HexColour117 Color = "#87d7ff"
// 	HexColour118 Color = "#87ff00"
// 	HexColour119 Color = "#87ff5f"
// 	HexColour120 Color = "#87ff87"
// 	HexColour121 Color = "#87ffaf"
// 	HexColour122 Color = "#87ffd7"
// 	HexColour123 Color = "#87ffff"
// 	HexColour124 Color = "#af0000"
// 	HexColour125 Color = "#af005f"
// 	HexColour126 Color = "#af0087"
// 	HexColour127 Color = "#af00af"
// 	HexColour128 Color = "#af00d7"
// 	HexColour129 Color = "#af00ff"
// 	HexColour130 Color = "#af5f00"
// 	HexColour131 Color = "#af5f5f"
// 	HexColour132 Color = "#af5f87"
// 	HexColour133 Color = "#af5faf"
// 	HexColour134 Color = "#af5fd7"
// 	HexColour135 Color = "#af5fff"
// 	HexColour136 Color = "#af8700"
// 	HexColour137 Color = "#af875f"
// 	HexColour138 Color = "#af8787"
// 	HexColour139 Color = "#af87af"
// 	HexColour140 Color = "#af87d7"
// 	HexColour141 Color = "#af87ff"
// 	HexColour142 Color = "#afaf00"
// 	HexColour143 Color = "#afaf5f"
// 	HexColour144 Color = "#afaf87"
// 	HexColour145 Color = "#afafaf"
// 	HexColour146 Color = "#afafd7"
// 	HexColour147 Color = "#afafff"
// 	HexColour148 Color = "#afd700"
// 	HexColour149 Color = "#afd75f"
// 	HexColour150 Color = "#afd787"
// 	HexColour151 Color = "#afd7af"
// 	HexColour152 Color = "#afd7d7"
// 	HexColour153 Color = "#afd7ff"
// 	HexColour154 Color = "#afff00"
// 	HexColour155 Color = "#afff5f"
// 	HexColour156 Color = "#afff87"
// 	HexColour157 Color = "#afffaf"
// 	HexColour158 Color = "#afffd7"
// 	HexColour159 Color = "#afffff"
// 	HexColour160 Color = "#d70000"
// 	HexColour161 Color = "#d7005f"
// 	HexColour162 Color = "#d70087"
// 	HexColour163 Color = "#d700af"
// 	HexColour164 Color = "#d700d7"
// 	HexColour165 Color = "#d700ff"
// 	HexColour166 Color = "#d75f00"
// 	HexColour167 Color = "#d75f5f"
// 	HexColour168 Color = "#d75f87"
// 	HexColour169 Color = "#d75faf"
// 	HexColour170 Color = "#d75fd7"
// 	HexColour171 Color = "#d75fff"
// 	HexColour172 Color = "#d78700"
// 	HexColour173 Color = "#d7875f"
// 	HexColour174 Color = "#d78787"
// 	HexColour175 Color = "#d787af"
// 	HexColour176 Color = "#d787d7"
// 	HexColour177 Color = "#d787ff"
// 	HexColour178 Color = "#d7af00"
// 	HexColour179 Color = "#d7af5f"
// 	HexColour180 Color = "#d7af87"
// 	HexColour181 Color = "#d7afaf"
// 	HexColour182 Color = "#d7afd7"
// 	HexColour183 Color = "#d7afff"
// 	HexColour184 Color = "#d7d700"
// 	HexColour185 Color = "#d7d75f"
// 	HexColour186 Color = "#d7d787"
// 	HexColour187 Color = "#d7d7af"
// 	HexColour188 Color = "#d7d7d7"
// 	HexColour189 Color = "#d7d7ff"
// 	HexColour190 Color = "#d7ff00"
// 	HexColour191 Color = "#d7ff5f"
// 	HexColour192 Color = "#d7ff87"
// 	HexColour193 Color = "#d7ffaf"
// 	HexColour194 Color = "#d7ffd7"
// 	HexColour195 Color = "#d7ffff"
// 	HexColour196 Color = "#ff0000"
// 	HexColour197 Color = "#ff005f"
// 	HexColour198 Color = "#ff0087"
// 	HexColour199 Color = "#ff00af"
// 	HexColour200 Color = "#ff00d7"
// 	HexColour201 Color = "#ff00ff"
// 	HexColour202 Color = "#ff5f00"
// 	HexColour203 Color = "#ff5f5f"
// 	HexColour204 Color = "#ff5f87"
// 	HexColour205 Color = "#ff5faf"
// 	HexColour206 Color = "#ff5fd7"
// 	HexColour207 Color = "#ff5fff"
// 	HexColour208 Color = "#ff8700"
// 	HexColour209 Color = "#ff875f"
// 	HexColour210 Color = "#ff8787"
// 	HexColour211 Color = "#ff87af"
// 	HexColour212 Color = "#ff87d7"
// 	HexColour213 Color = "#ff87ff"
// 	HexColour214 Color = "#ffaf00"
// 	HexColour215 Color = "#ffaf5f"
// 	HexColour216 Color = "#ffaf87"
// 	HexColour217 Color = "#ffafaf"
// 	HexColour218 Color = "#ffafd7"
// 	HexColour219 Color = "#ffafff"
// 	HexColour220 Color = "#ffd700"
// 	HexColour221 Color = "#ffd75f"
// 	HexColour222 Color = "#ffd787"
// 	HexColour223 Color = "#ffd7af"
// 	HexColour224 Color = "#ffd7d7"
// 	HexColour225 Color = "#ffd7ff"
// 	HexColour226 Color = "#ffff00"
// 	HexColour227 Color = "#ffff5f"
// 	HexColour228 Color = "#ffff87"
// 	HexColour229 Color = "#ffffaf"
// 	HexColour230 Color = "#ffffd7"
// 	HexColour231 Color = "#ffffff"
// 	HexColour232 Color = "#080808"
// 	HexColour233 Color = "#121212"
// 	HexColour234 Color = "#1c1c1c"
// 	HexColour235 Color = "#262626"
// 	HexColour236 Color = "#303030"
// 	HexColour237 Color = "#3a3a3a"
// 	HexColour238 Color = "#444444"
// 	HexColour239 Color = "#4e4e4e"
// 	HexColour240 Color = "#585858"
// 	HexColour241 Color = "#626262"
// 	HexColour242 Color = "#6c6c6c"
// 	HexColour243 Color = "#767676"
// 	HexColour244 Color = "#808080"
// 	HexColour245 Color = "#8a8a8a"
// 	HexColour246 Color = "#949494"
// 	HexColour247 Color = "#9e9e9e"
// 	HexColour248 Color = "#a8a8a8"
// 	HexColour249 Color = "#b2b2b2"
// 	HexColour250 Color = "#bcbcbc"
// 	HexColour251 Color = "#c6c6c6"
// 	HexColour252 Color = "#d0d0d0"
// 	HexColour253 Color = "#dadada"
// 	HexColour254 Color = "#e4e4e4"
// 	HexColour255 Color = "#eeeeee"
// ) // }}}

// // HTML Color Names
// const ( // {{{
// 	AliceBlue            Color = "#F0F8FF"
// 	AntiqueWhite         Color = "#FAEBD7"
// 	Aqua                 Color = "#00FFFF"
// 	Aquamarine           Color = "#7FFFD4"
// 	Azure                Color = "#F0FFFF"
// 	Beige                Color = "#F5F5DC"
// 	Bisque               Color = "#FFE4C4"
// 	Black                Color = "#000000"
// 	BlanchedAlmond       Color = "#FFEBCD"
// 	Blue                 Color = "#0000FF"
// 	BlueViolet           Color = "#8A2BE2"
// 	Brown                Color = "#A52A2A"
// 	BurlyWood            Color = "#DEB887"
// 	CadetBlue            Color = "#5F9EA0"
// 	Chartreuse           Color = "#7FFF00"
// 	Chocolate            Color = "#D2691E"
// 	Coral                Color = "#FF7F50"
// 	CornflowerBlue       Color = "#6495ED"
// 	Cornsilk             Color = "#FFF8DC"
// 	Crimson              Color = "#DC143C"
// 	Cyan                 Color = "#00FFFF"
// 	DarkBlue             Color = "#00008B"
// 	DarkCyan             Color = "#008B8B"
// 	DarkGoldenRod        Color = "#B8860B"
// 	DarkGray             Color = "#A9A9A9"
// 	DarkGrey             Color = "#A9A9A9"
// 	DarkGreen            Color = "#006400"
// 	DarkKhaki            Color = "#BDB76B"
// 	DarkMagenta          Color = "#8B008B"
// 	DarkOliveGreen       Color = "#556B2F"
// 	DarkOrange           Color = "#FF8C00"
// 	DarkOrchid           Color = "#9932CC"
// 	DarkRed              Color = "#8B0000"
// 	DarkSalmon           Color = "#E9967A"
// 	DarkSeaGreen         Color = "#8FBC8F"
// 	DarkSlateBlue        Color = "#483D8B"
// 	DarkSlateGray        Color = "#2F4F4F"
// 	DarkSlateGrey        Color = "#2F4F4F"
// 	DarkTurquoise        Color = "#00CED1"
// 	DarkViolet           Color = "#9400D3"
// 	DeepPink             Color = "#FF1493"
// 	DeepSkyBlue          Color = "#00BFFF"
// 	DimGray              Color = "#696969"
// 	DimGrey              Color = "#696969"
// 	DodgerBlue           Color = "#1E90FF"
// 	FireBrick            Color = "#B22222"
// 	FloralWhite          Color = "#FFFAF0"
// 	ForestGreen          Color = "#228B22"
// 	Fuchsia              Color = "#FF00FF"
// 	Gainsboro            Color = "#DCDCDC"
// 	GhostWhite           Color = "#F8F8FF"
// 	Gold                 Color = "#FFD700"
// 	GoldenRod            Color = "#DAA520"
// 	Gray                 Color = "#808080"
// 	Grey                 Color = "#808080"
// 	Green                Color = "#008000"
// 	GreenYellow          Color = "#ADFF2F"
// 	HoneyDew             Color = "#F0FFF0"
// 	HotPink              Color = "#FF69B4"
// 	IndianRed            Color = "#CD5C5C"
// 	Indigo               Color = "#4B0082"
// 	Ivory                Color = "#FFFFF0"
// 	Khaki                Color = "#F0E68C"
// 	Lavender             Color = "#E6E6FA"
// 	LavenderBlush        Color = "#FFF0F5"
// 	LawnGreen            Color = "#7CFC00"
// 	LemonChiffon         Color = "#FFFACD"
// 	LightBlue            Color = "#ADD8E6"
// 	LightCoral           Color = "#F08080"
// 	LightCyan            Color = "#E0FFFF"
// 	LightGoldenRodYellow Color = "#FAFAD2"
// 	LightGray            Color = "#D3D3D3"
// 	LightGrey            Color = "#D3D3D3"
// 	LightGreen           Color = "#90EE90"
// 	LightPink            Color = "#FFB6C1"
// 	LightSalmon          Color = "#FFA07A"
// 	LightSeaGreen        Color = "#20B2AA"
// 	LightSkyBlue         Color = "#87CEFA"
// 	LightSlateGray       Color = "#778899"
// 	LightSlateGrey       Color = "#778899"
// 	LightSteelBlue       Color = "#B0C4DE"
// 	LightYellow          Color = "#FFFFE0"
// 	Lime                 Color = "#00FF00"
// 	LimeGreen            Color = "#32CD32"
// 	Linen                Color = "#FAF0E6"
// 	Magenta              Color = "#FF00FF"
// 	Maroon               Color = "#800000"
// 	MediumAquaMarine     Color = "#66CDAA"
// 	MediumBlue           Color = "#0000CD"
// 	MediumOrchid         Color = "#BA55D3"
// 	MediumPurple         Color = "#9370DB"
// 	MediumSeaGreen       Color = "#3CB371"
// 	MediumSlateBlue      Color = "#7B68EE"
// 	MediumSpringGreen    Color = "#00FA9A"
// 	MediumTurquoise      Color = "#48D1CC"
// 	MediumVioletRed      Color = "#C71585"
// 	MidnightBlue         Color = "#191970"
// 	MintCream            Color = "#F5FFFA"
// 	MistyRose            Color = "#FFE4E1"
// 	Moccasin             Color = "#FFE4B5"
// 	NavajoWhite          Color = "#FFDEAD"
// 	Navy                 Color = "#000080"
// 	OldLace              Color = "#FDF5E6"
// 	Olive                Color = "#808000"
// 	OliveDrab            Color = "#6B8E23"
// 	Orange               Color = "#FFA500"
// 	OrangeRed            Color = "#FF4500"
// 	Orchid               Color = "#DA70D6"
// 	PaleGoldenRod        Color = "#EEE8AA"
// 	PaleGreen            Color = "#98FB98"
// 	PaleTurquoise        Color = "#AFEEEE"
// 	PaleVioletRed        Color = "#DB7093"
// 	PapayaWhip           Color = "#FFEFD5"
// 	PeachPuff            Color = "#FFDAB9"
// 	Peru                 Color = "#CD853F"
// 	Pink                 Color = "#FFC0CB"
// 	Plum                 Color = "#DDA0DD"
// 	PowderBlue           Color = "#B0E0E6"
// 	Purple               Color = "#800080"
// 	RebeccaPurple        Color = "#663399"
// 	Red                  Color = "#FF0000"
// 	RosyBrown            Color = "#BC8F8F"
// 	RoyalBlue            Color = "#4169E1"
// 	SaddleBrown          Color = "#8B4513"
// 	Salmon               Color = "#FA8072"
// 	SandyBrown           Color = "#F4A460"
// 	SeaGreen             Color = "#2E8B57"
// 	SeaShell             Color = "#FFF5EE"
// 	Sienna               Color = "#A0522D"
// 	Silver               Color = "#C0C0C0"
// 	SkyBlue              Color = "#87CEEB"
// 	SlateBlue            Color = "#6A5ACD"
// 	SlateGray            Color = "#708090"
// 	SlateGrey            Color = "#708090"
// 	Snow                 Color = "#FFFAFA"
// 	SpringGreen          Color = "#00FF7F"
// 	SteelBlue            Color = "#4682B4"
// 	Tan                  Color = "#D2B48C"
// 	Teal                 Color = "#008080"
// 	Thistle              Color = "#D8BFD8"
// 	Tomato               Color = "#FF6347"
// 	Turquoise            Color = "#40E0D0"
// 	Violet               Color = "#EE82EE"
// 	Wheat                Color = "#F5DEB3"
// 	White                Color = "#FFFFFF"
// 	WhiteSmoke           Color = "#F5F5F5"
// 	Yellow               Color = "#FFFF00"
// 	YellowGreen          Color = "#9ACD32"
// ) // }}}

func ColorOn() {
	showColor = true
}

func ColorOff() {
	showColor = false
}

func IsColor() bool {
	return showColor
}

// TernaryIf es una función que devuelve un valor dependiendo de una condición, similar al operador ternario `?:`
// para que funcione `T` debe ser del mismo tipo de dato.
// Esta función se usa con frecuencia como atajo para la instrucción if.
//
// Parámetros:
// - `condition`: Una expresión que se evalúa como true o false.
// - `trueVal`, `falseVal`: Expresion.o con valores de algún tipo.
// Resultado:
// - Si `condition` es true, el operador retorna el valor de `trueVal`; de lo contrario, devuelve el valor de `falseVal`.
func ternaryIf[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}

	return falseVal
}