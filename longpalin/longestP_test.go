package longpalin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromeV1(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"kajsnoiewpwerjfpwefnkjcaacaacaacj", "jcaacaacaacj"},
		{"bahgfdabababahgfdbad", "abababa"},
		{"babad", "aba"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
		{"acc", "cc"},
		{"acac", "aca"},
	}

	for _, v := range tests {
		t.Log("Input: ", v.input)
		got := longestPalindromeV1(v.input)
		assert.Equal(t, v.want, got)
	}
}

func TestLongestPalindromeV2(t *testing.T) {
	tests := []struct {
		input string
		want1 string
		want2 string
	}{
		{"aaaa", "aaaa", ""},
		{"eeeee", "eeeee", ""},
		{"ccc", "ccc", ""},
		{"ffffff", "ffffff", ""},
		{"bb", "bb", ""},
		{"ccd", "cc", ""},
		{
			"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth",
			"ranynar",
			"",
		},
		{
			"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearthtraeehtmorfhsireptonllahselpoepehtrofelpoepehtybelpoepehtfotnemnrevogtahtdnamodeerffohtribwenaevahllahsdoGredsnunoitansihttahtniavnideidevahtonllahsdaedesehttahtevloserylhgiherehewtahtnoitovedfoerusaemllufptsalehtevagyehthcihwrofesuactahtotnoitoveddesaercniekatewdaedderonohesehtmorftahtsuerofebgniniamerksfadttaergehtotdetacidederehebotsurofrehtarsitIdecnavdaylbonosrafsuhtevaherehthguofohwyehthcihwkrowdehsinifnluehtoterehdetacidedebotrehtargnivilehtsurofsitIerehdidyehttahwtegrofrevennactituberehyasewtahwrebmemergnolroneltonelttilllifwsdadlrowehgTtcartedroddaotrewnoproopruoevobaraftidetarcesnocevaherehdelggurtsohwdaeddnagnivilnemlevarbehTdnuorgsihtwollahtonnacewetarcesnoctonnacewetacidedtonnacewesnesregralanituBsihtoddluohsewtahtreporpdnagnafrehtegotlasitIevilthgimnoitantahttahtsevilriehtevagerehohwesohtrofecalpgnitserlanifasadleiftahtfonoitropaetapcidedotemocevaheWrawtahztfodlmeifelttabtaerganotemqeraeWerudnegnolnacdetacidedosdnadeviecnocosnoitranynaronoitpantahtrehtehwgnitsetrawlivic",
			"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearthtraeehtmorfhsireptonllahselpoepehtrofelpoepehtybelpoepehtfotnemnrevogtahtdnamodeerffohtribwenaevahllahsdoGredsnunoitansihttahtniavnideidevahtonllahsdaedesehttahtevloserylhgiherehewtahtnoitovedfoerusaemllufptsalehtevagyehthcihwrofesuactahtotnoitoveddesaercniekatewdaedderonohesehtmorftahtsuerofebgniniamerksfadttaergehtotdetacidederehebotsurofrehtarsitIdecnavdaylbonosrafsuhtevaherehthguofohwyehthcihwkrowdehsinifnluehtoterehdetacidedebotrehtargnivilehtsurofsitIerehdidyehttahwtegrofrevennactituberehyasewtahwrebmemergnolroneltonelttilllifwsdadlrowehgTtcartedroddaotrewnoproopruoevobaraftidetarcesnocevaherehdelggurtsohwdaeddnagnivilnemlevarbehTdnuorgsihtwollahtonnacewetarcesnoctonnacewetacidedtonnacewesnesregralanituBsihtoddluohsewtahtreporpdnagnafrehtegotlasitIevilthgimnoitantahttahtsevilriehtevagerehohwesohtrofecalpgnitserlanifasadleiftahtfonoitropaetapcidedotemocevaheWrawtahztfodlmeifelttabtaerganotemqeraeWerudnegnolnacdetacidedosdnadeviecnocosnoitranynaronoitpantahtrehtehwgnitsetrawlivic",
			"",
		},
		{"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldlmeifelttabtaerganotemqeraeWerudnegnolnacdetacidedosdnadeviecnocosnoitranynaronoitpantahtrehtehwgnitsetrawlivic", "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldlmeifelttabtaerganotemqeraeWerudnegnolnacdetacidedosdnadeviecnocosnoitranynaronoitpantahtrehtehwgnitsetrawlivic", ""},
		{"kajsnoiewpwerjfpwefnkjcaacaacaacj", "jcaacaacaacj", ""},
		{"bahgfdabababahgfdbad", "abababa", ""},
		{"babad", "bab", "aba"},
		{"cbbd", "bb", ""},
		{"a", "a", ""},
		{"ac", "a", "c"},
		{"acc", "cc", ""},
		{"acac", "aca", "cac"},
	}

	for _, v := range tests {
		t.Log("Input: ", v.input)
		got := longestPalindromeV2(v.input)
		if got != v.want1 && got != v.want2 {
			t.Errorf("ERROR: Wanted: %s or %s; got %s", v.want1, v.want2, got)
		}
	}
}

func TestLongestPalindromeV3(t *testing.T) {
	tests := []struct {
		input string
		want1 string
		want2 string
	}{
		{"adam", "ada", ""},
		{"adama", "ada", "ama"},
		{"adada", "adada", ""},
		{"aaaa", "aaaa", ""},
		{"aaaab", "aaaa", ""},
		{"baaaa", "aaaa", ""},
		{"eeeee", "eeeee", ""},
		{"ccc", "ccc", ""},
		{"ffffff", "ffffff", ""},
		{"bb", "bb", ""},
		{"ccd", "cc", ""},
		{"cdd", "dd", ""},
		{
			"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth",
			"ranynar",
			"",
		},
		{
			"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearthtraeehtmorfhsireptonllahselpoepehtrofelpoepehtybelpoepehtfotnemnrevogtahtdnamodeerffohtribwenaevahllahsdoGredsnunoitansihttahtniavnideidevahtonllahsdaedesehttahtevloserylhgiherehewtahtnoitovedfoerusaemllufptsalehtevagyehthcihwrofesuactahtotnoitoveddesaercniekatewdaedderonohesehtmorftahtsuerofebgniniamerksfadttaergehtotdetacidederehebotsurofrehtarsitIdecnavdaylbonosrafsuhtevaherehthguofohwyehthcihwkrowdehsinifnluehtoterehdetacidedebotrehtargnivilehtsurofsitIerehdidyehttahwtegrofrevennactituberehyasewtahwrebmemergnolroneltonelttilllifwsdadlrowehgTtcartedroddaotrewnoproopruoevobaraftidetarcesnocevaherehdelggurtsohwdaeddnagnivilnemlevarbehTdnuorgsihtwollahtonnacewetarcesnoctonnacewetacidedtonnacewesnesregralanituBsihtoddluohsewtahtreporpdnagnafrehtegotlasitIevilthgimnoitantahttahtsevilriehtevagerehohwesohtrofecalpgnitserlanifasadleiftahtfonoitropaetapcidedotemocevaheWrawtahztfodlmeifelttabtaerganotemqeraeWerudnegnolnacdetacidedosdnadeviecnocosnoitranynaronoitpantahtrehtehwgnitsetrawlivic",
			"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearthtraeehtmorfhsireptonllahselpoepehtrofelpoepehtybelpoepehtfotnemnrevogtahtdnamodeerffohtribwenaevahllahsdoGredsnunoitansihttahtniavnideidevahtonllahsdaedesehttahtevloserylhgiherehewtahtnoitovedfoerusaemllufptsalehtevagyehthcihwrofesuactahtotnoitoveddesaercniekatewdaedderonohesehtmorftahtsuerofebgniniamerksfadttaergehtotdetacidederehebotsurofrehtarsitIdecnavdaylbonosrafsuhtevaherehthguofohwyehthcihwkrowdehsinifnluehtoterehdetacidedebotrehtargnivilehtsurofsitIerehdidyehttahwtegrofrevennactituberehyasewtahwrebmemergnolroneltonelttilllifwsdadlrowehgTtcartedroddaotrewnoproopruoevobaraftidetarcesnocevaherehdelggurtsohwdaeddnagnivilnemlevarbehTdnuorgsihtwollahtonnacewetarcesnoctonnacewetacidedtonnacewesnesregralanituBsihtoddluohsewtahtreporpdnagnafrehtegotlasitIevilthgimnoitantahttahtsevilriehtevagerehohwesohtrofecalpgnitserlanifasadleiftahtfonoitropaetapcidedotemocevaheWrawtahztfodlmeifelttabtaerganotemqeraeWerudnegnolnacdetacidedosdnadeviecnocosnoitranynaronoitpantahtrehtehwgnitsetrawlivic",
			"",
		},
		{"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldlmeifelttabtaerganotemqeraeWerudnegnolnacdetacidedosdnadeviecnocosnoitranynaronoitpantahtrehtehwgnitsetrawlivic", "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldlmeifelttabtaerganotemqeraeWerudnegnolnacdetacidedosdnadeviecnocosnoitranynaronoitpantahtrehtehwgnitsetrawlivic", ""},
		{"kajsnoiewpwerjfpwefnkjcaacaacaacj", "jcaacaacaacj", ""},
		{"bahgfdabababahgfdbad", "abababa", ""},
		{"babad", "bab", "aba"},
		{"cbbd", "bb", ""},
		{"a", "a", ""},
		{"ac", "a", "c"},
		{"acc", "cc", ""},
		{"abc", "a", "b"},
		{"acac", "aca", "cac"},
	}

	for _, v := range tests {
		t.Log("Input: ", v.input)
		got := longestPalindromeV3(v.input)
		if got != v.want1 && got != v.want2 {
			t.Errorf("ERROR: Wanted: %s or %s; got %s", v.want1, v.want2, got)
		}
	}
}
