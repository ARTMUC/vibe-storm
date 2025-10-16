# Frontend Refactoring Plan

## Current State Analysis

### Issues Identified:
1. **Overuse of React**: Entire page is rendered as React SPA when most content is static
2. **Code Duplication**: `AuthStatus.tsx` and `APITesting.tsx` exist separately but are duplicated inside `App.tsx`
3. **Poor Separation**: No clear boundary between server-rendered (Templ) and client-side (React) content
4. **Monolithic Component**: `App.tsx` is 500+ lines containing everything from static hero sections to complex forms

### What Should Be in Templ (Server-Rendered):
- ✅ Base layout (header, footer, navigation) - Already done
- ❌ Hero section - Currently in React, should be Templ
- ❌ Features grid section - Currently in React, should be Templ  
- ❌ API status indicator - Currently in React, could be Templ + minimal JS
- ❌ Section structure and layout - Currently in React, should be Templ

### What Should Stay in React (Complex Interactivity):
- ✅ Authentication status widget (token management, refresh, localStorage interaction)
- ✅ API testing interface (tabbed forms, async operations, complex state)

## Refactoring Strategy

### Phase 1: Create Templ Components
1. Create `web/templates/components/hero.templ` - Static hero section
2. Create `web/templates/components/features.templ` - Static features grid
3. Update `web/templates/home.templ` - Compose page with Templ components + React mount points

### Phase 2: Refactor React Components
1. **Keep**: `AuthStatus.tsx` - Focused component for auth management
2. **Keep**: `APITesting.tsx` - Focused component for API testing
3. **Delete**: `App.tsx` - Monolithic component, functionality split between Templ and focused React components
4. **Update**: `index.tsx` - Mount multiple React components at specific DOM nodes

### Phase 3: Update Build & Integration
1. Update webpack configuration if needed
2. Ensure React components mount correctly in Templ-rendered pages
3. Test all functionality

## Benefits of This Approach

### Performance:
- Faster initial page load (less JS to parse/execute)
- Better SEO (static content rendered server-side)
- Reduced bundle size (only interactive parts in React)

### Maintainability:
- Clear separation of concerns
- Static content easy to update in Templ
- React used only where necessary
- No code duplication

### Best Practices:
- Server-render what you can, hydrate what you must
- Progressive enhancement approach
- Follows Go + Templ ecosystem patterns

## Implementation Order

1. ✅ Review existing code (COMPLETED)
2. Create Templ component templates
3. Update home.templ with new structure
4. Simplify React components (remove App.tsx)
5. Update index.tsx entry point
6. Test and verify functionality
