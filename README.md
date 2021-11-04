Ruin Generator
==============

This is my attempt to build a program to generate ruins for Numenera using the
rules from the Jade Colossus splatbook.

The output only contains the roll results from the ruin generator though, none
of the text.

## How This Works

This program generates a ruin using the [Numenera Ruin Mapping
Engine](https://www.montecookgames.com/store/product/jade-colossus-ruins-of-the-prior-worlds/)
(RME).

It makes a few assumptions right off the bat:

* there is a single entrance into the ruin 
* ruins don't go deeper than 10 levels

The reasons for those are pretty straightforward. Having a single entrance makes
the code a lot easier to manage, as we're just building a tree data
structure. Not going deeper than 10 levels is so that the ruins don't
accidentally become massive sprawling affairs. 


# Features 

There are 15 types of features that can be rolled in the RME:

1. Corridor 
2. Chamber
3. Creature 
4. Explorers 
5. Interstitial cavity 
6. Accessway 
7. Rupture 
8. Shaft 
9. Abhuman colony 
10. Integrated machine 
11. Matter leak
12. Energy discharge 
13. Weird event 
14. Vault
15. Relic chamber 

However, not all of these results create new rooms or passageways -- some modify
the previously created chamber.

As such, there are three 'types' of feature:

* rooms ( chambers, interstitial cavities, etc )
* passageways ( corridors, accessways, shafts )
* modifiers ( weird events, etc )

However, as the only difference between a "room" and a "passageway" is shape,
size, and function we'll only refer to "rooms" and "modifiers". The reason for
this is that an empty chamber ( especially ones on the smaller side of things )
are functionally identical to a corridor. 

## Rooms 

These are features that describe a room ( although interstitial cavities make
"room" do a lot of work there ). Each room always has at least one exit -- the
one that brought the adventurers into that room.

There are two "types" of room: 'base' room types and 'fancy' room types.

**Base room types**

1. Corridor
2. Accessway
3. Shaft
4. Chamber (15% have no defining feature or are otherwise empty)
5. Interstitial cavity (+1d10 additional exits)
6. Vault
7. Relic chamber 

**Fancy room types**

1. Creature
3. Explorers
5. Abhuman colony
6. Integrated machine
7. Matter leak
8. Energy discharge
9. Rupture


"Fancy" rooms are ones that when you roll that result, you usually then make a
secondary roll to determine what kind of base room the feature is contained
within. For example, if a creature is rolled on the main feature table then the
room it ends up in could be a corridor, chamber, interstitial cavity, shaft, or
a rupture.

Additionally, rooms can have additional rules about what happens when they are
generated. These include the fact that relic chambers have an 80% chance of
containing a creature, or that the "chamber" room type has a 15% chance of being
completely empty.

## Modifiers

There's really only one "modifier" type of main feature: weird events.

These get added to the list of features for the previously generated room.


==============================

ruin creation process 

entrance
 -> roll on corridor table ( this corridor can't be modified, rolling a rupture or weird event causes a re-roll when rolling on this rooms exits )
   -> roll on exit table ( min 1 for this roll )
     -> for each exit:
       -> roll on main feature table 
         -> result is a modifier?
           => yes
             -> add to 'modifier' list of previous room, roll again 
           => no
             -> connect the new room and the previous one via the exit
          



Entrance ()
  
becomes (for example):

Entrance (
  Corridor (type: 4, exits: [regular])
)  









Start: 

entrance -> corridor (roll) + exits (roll)




Node -> Node
Node: Main Feature ( corridor, chamber, shaft, rupture, etc )

