package qoutes

import "math/rand"

var collection = []string{
	"“Absorb what is useful, Discard what is not, add what is uniquely your own.” — Bruce Lee",
	"“Choose the positive. You have choice, you are master of your attitude, choose the positive, the constructive. Optimism is a faith that leads to success.” — Bruce Lee",
	"“Don’t fear failure.  Not failure, but low aim, is the crime. In great attempts it is glorious even to fail.” — Bruce Lee",
	"“Empty your mind; be formless, shapeless – like water. Now you put water into a cup, it becomes the cup, you put water into a bottle, it becomes the bottle, you put it in a teapot, it becomes the teapot. Now water can flow, or it can crash. Be water, my friend.” — Bruce Lee",
	"“I fear not the man who has practiced 10,000 kicks once, but I fear the man who has practiced one kick 10,000 times.” — Bruce Lee",
	"“It’s not the daily increase but daily decrease. Hack away at the unessential.” — Bruce Lee",
	"“Mistakes are always forgivable, if one has the courage to admit them.” — Bruce Lee",
	"“To hell with circumstances; I create opportunities.” — Bruce Lee",
	"“What is” is more important than ‘what should be.’ Too many people are looking at ‘what is’ from a position of thinking ‘what should be’.” — Bruce Lee",
	"“When one has reached maturity in the art, one will have a formless form. It is like ice dissolving in water. When one has no form, one can be all forms; when one has no style, he can fit in with any style.” — Bruce Lee",
}

func RandomQuote() string {
	return collection[rand.Intn(len(collection))]
}
