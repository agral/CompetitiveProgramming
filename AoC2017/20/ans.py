#!/usr/bin/env python3

"""
Solution to the challenge #20 of the "Advent of Code 2017" series.

MIT License

Copyright (c) 2017 Adam Szczerbiak (szczerbiakadam@gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"""

import math

class Vector3:
    def __init__(self, x, y, z):
        self.x = x
        self.y = y
        self.z = z

    def __add__(self, vector3):
        self.x += vector3.x
        self.y += vector3.y
        self.z += vector3.z
        return self

    def __str__(self):
        return "<{},{},{}>".format(self.x, self.y, self.z)

    def magnitude(self):
        return math.sqrt(self.x ** 2 + self.y ** 2 + self.z ** 2)

    def manhattan(self):
        return abs(self.x) + abs(self.y) + abs(self.z)

class Particle3:
    def __init__(self, position, velocity, acceleration):
        self.position = position
        self.velocity = velocity
        self.acceleration = acceleration
        self.has_collided = False

    def __str__(self):
        return "p={}, v={}, a={}".format(
                self.position, self.velocity, self.acceleration)

    def step(self):
        self.velocity += self.acceleration
        self.position += self.velocity


data_file = "input_20"

def read_data():
    def string_to_vector3(string):
        """ Converts a string such as "p=<0,1,-2>" to a Vector3(0,1,-2)."""
        data = string[3:-1] # trims out the useless "p=<" and ">" characters.
        separated = data.split(",")
        return Vector3(int(separated[0]), int(separated[1]), int(separated[2]))

    particles_collection = []
    data_lines = [line.strip() for line in open(data_file, "r")]
    for line_of_data in data_lines:
        serialized_vectors = line_of_data.split(", ")
        position = string_to_vector3(serialized_vectors[0])
        velocity = string_to_vector3(serialized_vectors[1])
        acceleration = string_to_vector3(serialized_vectors[2])

        particles_collection.append(Particle3(position, velocity, acceleration))
    print("Loaded {} particles.".format(len(particles_collection)))
    return particles_collection


# Challenge A: which particle in the long-run remains closest to the origin?
LONG_RUN = 50000
particles = read_data()

for r in range(LONG_RUN):
    text = "Calculating step {}/{} ({:.2f}%)...\r".format(
            r+1, LONG_RUN, 100 * (r+1) / LONG_RUN)
    print(text, end="")

    positions = set()
    collisions = set()

    for particle in particles:
        particle.step()
        if str(particle.position) in positions:
            collisions.add(str(particle.position))
        if not particle.has_collided:
            positions.add(str(particle.position))

    for particle in particles:
        if str(particle.position) in collisions:
            particle.has_collided = True


print("\nDone!")

best_particle = 0
best_distance = particles[0].position.manhattan()
for p in range(len(particles)):
    d = particles[p].position.manhattan()
    if d < best_distance:
        best_distance = d
        best_particle = p
print("Challenge A: #{} wins with manhattan({}).".format(
    best_particle, best_distance))

noncollided_particles = [p for p in particles if not p.has_collided]
print("Challenge B: {} particles have not collided.".format(
        len(noncollided_particles)))
