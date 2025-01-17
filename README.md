# gone

[![Github Actions Widget]][github actions] [![GoReport Widget]][goreport] [![GoDoc Widget]][godoc]

A simple neural network library in Go from scratch. 0 dependencies\*

_there are 0 neural network related dependencies, the only dependencies are for testing ([stretchr/testify](github.com/stretchr/testify)) and for persistence ([golang/protobuf](github.com/golang/protobuf))_

[goreport widget]: https://goreportcard.com/badge/github.com/fr3fou/gone
[goreport]: https://goreportcard.com/report/github.com/fr3fou/gone
[github actions widget]: https://github.com/fr3fou/gone/workflows/Test/badge.svg
[github actions]: https://github.com/fr3fou/gone/actions
[godoc]: http://pkg.go.dev/github.com/fr3fou/gone
[godoc widget]: https://godoc.org/github.com/fr3fou/gone?status.svg

## Example

### Getting started

```go
	g := gone.New(
		0.1,
		gone.MSE(),
		gone.Layer{
			Nodes: 2,
		},
		gone.Layer{
			Nodes:     4,
			Activator: gone.Sigmoid(),
		},
		gone.Layer{
			Nodes: 1,
		},
	)

	g.Train(gone.SGD(), gone.DataSet{
		{
			Inputs:  []float64{1, 0},
			Targets: []float64{1},
		},
		{
			Inputs:  []float64{0, 1},
			Targets: []float64{1},
		},
		{
			Inputs:  []float64{1, 1},
			Targets: []float64{0},
		},
		{
			Inputs:  []float64{0, 0},
			Targets: []float64{0},
		},
	}, 5000)

	g.Predict([]float64{1, 1})
```

### Saving model to disk

```go
	g.Save("test.gone")

```

### Loading model back into memory

```go
	g, err := gone.Load("test.gone")
```

## TODO

### `gone/`

- [ ] Types of task:
  - [x] Classification
  - [ ] Regression
- [x] Bias
  - [x] Matrix, rather than a single number
- [x] Feedforward (Predict)
- [ ] Train
  - [x] Support shuffling the data
  - [x] Epochs
  - [x] Backpropagation
  - [x] Batching
  - [ ] Different loss functions
    - [x] Mean Squared Error
    - [ ] Cross Entropy Error
- [x] Saving data - Done thanks to protobuf
- [x] Loading data
- [ ] Adam optimizer
- [ ] Nestrov + Momentum for GD
- [x] Fix MSE computation in debug mode (not used in actual backpropagation)
- [ ] Flatten layer
- [ ] Convolutional Layers


### `matrix/`

- [x] Randomize
- [x] Transpose
- [x] Scale
- [x] AddMatrix
- [x] Add
- [x] SubtractMatrix
- [x] Subtract
- [x] Multiply
- [x] Multiply
- [x] Flatten
- [x] Unflatten
- [x] NewFromArray - makes a single row
- [x] Map
- [x] Fold
- [x] Methods to support chaining

```go
	    n.Weights[i].
		Multiply(output).                         // weighted sum of the previous layer)
		Add(n.Layers[i+1].Bias).                  // bias
		Map(func(val float64, x, y int) float64 { // activation
			return n.Layers[i+1].Activator.F(val)
		})
```

### Research

- [x] Derivatives ~
- [x] Partial Derivatives ~
- [ ] Linear vs non-linear problems (activation function)
- [x] Gradient Descent
  - [x] (Batch) Gradient Descent (GD)
  - [x] Stochastic Gradient Descent (SGD)
  - [x] Mini-Batch Gradient Descent (MBGD?)
- [ ] Softmax (needed for multi class classification!)
- [x] Mean Squared Error
- [ ] Cross Entropy Error (needed for multi class classification!)
- [ ] How to determine how many layers and nodes to use
- [ ] One Hot Encoding

### Examples

- [x] XOR Problem
- [x] [Digit Classifier](https://github.com/fr3fou/digit-classifier)

### Shoutouts

- [David Josephs](https://github.com/josephsdavid) - was of HUGE help with algebra and other ML-related questions; also helped me spot some nasty bugs!

## Resources used:

- https://www.analyticsvidhya.com/blog/2020/01/fundamentals-deep-learning-activation-functions-when-to-use-them/
- https://www.youtube.com/watch?v=XJ7HLz9VYz0&list=PLRqwX-V7Uu6Y7MdSCaIfsxc561QI0U0Tb
- https://www.youtube.com/watch?v=aircAruvnKk&list=PLZHQObOWTQDNU6R1_67000Dx_ZCJB-3pi
- http://matrixmultiplication.xyz/
- https://www.khanacademy.org/math/precalculus/x9e81a4f98389efdf:matrices/x9e81a4f98389efdf:properties-of-matrix-addition-and-scalar-multiplication/a/properties-of-matrix-addition
- https://www.wikiwand.com/en/Matrix_(mathematics)
- https://www.wikiwand.com/en/Activation_function
- https://www.wikiwand.com/en/Delta_rule
- https://www.jeremyjordan.me/intro-to-neural-networks/
- https://www.arxiv-vanity.com/papers/2003.02139/
- https://machinelearningmastery.com/gentle-introduction-mini-batch-gradient-descent-configure-batch-size/
- http://neuralnetworksanddeeplearning.com/chap2.html
- https://arxiv.org/pdf/1802.01528.pdf
- https://github.com/stevenmiller888/mind/blob/master/index.js
- https://github.com/stevenmiller888/go-mind
- https://medium.com/yottabytes/everything-you-need-to-know-about-gradient-descent-applied-to-neural-networks-d70f85e0cc14
- https://towardsdatascience.com/deep-learning-which-loss-and-activation-functions-should-i-use-ac02f1c56aa8
- https://github.com/milosgajdos/go-neural
- https://ml-cheatsheet.readthedocs.io/en/latest/forwardpropagation.html
- https://mattmazur.com/2015/03/17/a-step-by-step-backpropagation-example/
- https://medium.com/coinmonks/representing-neural-network-with-vectors-and-matrices-c6b0e64db9fb
- https://towardsdatascience.com/classifying-cat-pics-with-a-logistic-regression-model-e35dfb9159bb
- https://towardsdatascience.com/math-neural-network-from-scratch-in-python-d6da9f29ce65
- https://towardsdatascience.com/gradient-descent-from-scratch-e8b75fa986cc
- https://rstudio-pubs-static.s3.amazonaws.com/337306_79a7966fad184532ab3ad66b322fe96e.html
