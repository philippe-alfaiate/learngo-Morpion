# Tuto Morpion

Apprendre à faire un morpion en go

## Préparation

Copier le dossier ACopier et le renommer en PrénomNom
puis ouvrir le dossier avec l'application vscode

## Commencer à programmer

Il y a une règle à suivre : Avant de modifier le document vérifier que c'est bien ce que l'on vous demande il y aura des *** pour prévenir lorsque vous devez modifier

Voici le contenu du fichier morpion.go

```go
package main

import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

var notClose bool

func main() {
	notClose = true
	for notClose {
		pixelgl.Run(run)
	}

}

func run() {

}

```

Que veux dire se code ?

La première ligne: 

```go
package main
```
Pour faire simple, c'est pour dire que notre application se lancera directement (un jeu, une application)


La ligne suivante:

```go
import (
	"fmt"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)
```
C'est pour dire quelle bibliothèque on utilise.
Mais c'est quoi une bibliothèque ?
C'est un ensemble de fonctions qui ont déjà été faite pour éviter de refaire des choses.
>	par exemple : ```math.max(5, 6)``` Cette fonction calcule le maximum entre 2 nombre, c'est plus rapide de l'utilier plutôt que de la refaire. Elle est contenu dans la bibliothèque ```"math"```




La ligne suivante:

```go
var notClose bool
```

```go
func main() {
	notClose = true
	for notClose {
		pixelgl.Run(run)
	}

}
```

```go
func run() {

}

```



* [Choisir une couleur](https://www.w3schools.com/colors/colors_picker.asp) - Couleur

##

A step by step series of examples that tell you how to get a development env running

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Billie Thompson** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc