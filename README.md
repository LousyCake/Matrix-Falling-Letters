# Matrix-Style Falling Letters Simulation

Simulate the iconic Matrix-style falling letters animation using Go (Golang) and the [gg](https://pkg.go.dev/github.com/fogleman/gg) graphics library.

## Overview

This Go program generates a mesmerizing simulation of characters falling like raindrops in the Matrix movie. Each character leaves a fading trail as it descends, creating a visually appealing effect.

## Features

- **Customizable Parameters**: Easily adjust simulation parameters such as the number of particles, fall speed range, trail length, and more.
- **Dynamic Character Changes**: Characters in the falling trail dynamically change every few frames, adding variety to the display.
- **Frame Saving**: Save each frame as a PNG file to analyze or create GIFs from the simulation.
- **Video Generation**: Combine saved frames into a video using [ffmpeg](https://ffmpeg.org/) to share the entire simulation.

## Dependencies

- [gg](https://pkg.go.dev/github.com/fogleman/gg): A 2D graphics library for Go.

## Installation

1. Install the required dependency:

    ```
    go get -u github.com/fogleman/gg
    ```

2. Clone the repository:

    ```
    git clone <repository_url>
    ```

3. Run the simulation:

    ```
    go run main.go
    ```

## Configuration

Adjust the simulation parameters by modifying the constants in the `main.go` file:

- `numParticles`: Number of falling particles.
- `minFallSpeed` and `maxFallSpeed`: Range of falling speeds.
- `fadeRate`: Rate at which the characters fade.
- `trailLength`: Length of the fading trail.
- `outputDir`: Directory to save simulation frames.

## Usage

1. Run the simulation and press 'Enter' to stop it.
2. Frames are saved in the specified `output` directory.

## Generating a Video

To combine the saved frames into a video, you can use a tool like [ffmpeg](https://ffmpeg.org/). Here is an example command:

```
ffmpeg -framerate 30 -i output/frame_%03d.png -c:v libx264 -r 30 -pix_fmt yuv420p matrix_animation.mp4
```
Adjust the input and output filenames as needed. This command assumes a frame rate of 30 frames per second, matching the simulation's frame rate. 

Now you should have a video file (matrix_animation.mp4) capturing the entire Matrix-style falling letters simulation.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

Feel free to customize the content further based on your project's specific details.
